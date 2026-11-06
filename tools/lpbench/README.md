# Log poller unfinalized-block batching benchmark (L1)

Status: **draft experiment**. This directory contains a read-only RPC probe and a
local mock server for deciding whether the L1 proposal in the Arbitrum VRF /
BHS incident RCA is worth implementing, and what chunk size it needs.

The probe does **not** modify the node, its database, or its RPC configuration.
It replays the exact two RPC calls the log poller makes for unfinalized blocks
over a fixed historical range, against whichever endpoints you give it.

## TL;DR

- L1 changes the unfinalized-block walk from two sequential RPC calls per block
  to two JSON-RPC batch calls per chunk. Whether that is a 1x or 10x win depends
  almost entirely on the provider executing batch elements in parallel.
- The existing metrics already show LinkPool is below Arbitrum's block rate:
  mean header + mean `FilterLogs` latency gives **~2.4 blocks/s**, versus
  Arbitrum's ~3.7 blocks/s. That matches the observed ~2.3 blocks/s.
- A control run against a public Arbitrum RPC measured **3.5 blocks/s** in the
  current per-block mode and **23.1 blocks/s** with a batch of 8 (6.6x). That is
  not a prediction for LinkPool, but it proves the effect is real and
  measurable.
- Start with chunk **8–16**, not 50. Chunk 50 is close to the 10 s RPC timeout
  in the serial case, and it widens the reorg window.
- Before implementing L1, fix the correctness gaps listed in
  [L1 implementation review notes](#l1-implementation-review-notes).

## 1. What L1 changes

Current loop (`pkg/logpoller/log_poller.go`, `getUnfinalizedLogs`):

```text
for each unfinalized block:
    eth_getLogs({blockHash: current.hash})     # 1 RTT
    eth_getBlockByNumber(current.number + 1)   # 1 RTT
    verify next.parentHash == current.hash
    current = next
```

Two sequential calls per block. At ~160 ms per call this caps throughput at
~3.1 blocks/s; a serial pipeline is limited by the mean latency, including the
tail, which is why the observed rate was ~2.3 blocks/s.

L1 proposes:

```text
for each chunk of N blocks:
    fetch N headers in one JSON-RPC batch       # 1 RTT
    verify the parent-hash chain locally
    fetch N logs by block hash in one batch    # 1 RTT
```

Two batch round trips per chunk. Throughput is:

```text
N / (B_headers(N) + B_logs(N))
```

If the provider processes the batch serially, `B(N) ~= N * latency`, and there
is no speedup. If it processes in parallel with effective concurrency `C`,
throughput approaches `C / (2 * latency)`. The probe measures `B(N)` directly.

One important detail: a single `eth_getLogs` filter has one `blockHash` field
(`pkg/client/compatibility_helper.go`). The log batch must be N independent
`eth_getLogs` requests in one JSON-RPC batch, or a range query. It cannot be one
filter with N hashes.

## 2. What the probe measures

For each endpoint and each batch size `N`:

- wall-clock time for the legacy per-block walk over a fixed range
- wall-clock time for the batched walk over the same range
- blocks/s and speedup versus legacy
- whether the returned log set is identical to legacy
- batch failures (`413`, `429`, JSON-RPC errors, timeouts)

The batch-size sweep is the important part. Plotting `B(N)` versus `N` tells you
whether the provider is:

- **parallel**: `B(N)` stays roughly flat, throughput scales with `N`
- **serial**: `B(N)` grows linearly, batching is just a wrapper around the same
  number of serial calls
- **capped**: `B(N)` flattens at some effective concurrency `C`, so batches
  larger than `C` add latency but no throughput

## 3. Quick start

### 3.1 Local validation (no credentials)

```bash
python3 tools/lpbench/mock_rpc_server.py \
  --port 8545 --delay-ms 20 --batch-concurrency 16 &

python3 tools/lpbench/lp_batch_probe.py \
  --rpc mock=http://127.0.0.1:8545 \
  --blocks 16 --batch-sizes 1,2,4,8,16 --runs 1
```

Use `--batch-concurrency 1` to simulate a provider that serializes batches.
Use `--max-batch 4` to simulate a provider batch-size limit.

### 3.2 Real endpoints

```bash
export ALCHEMY_URL=...    # from the k8s secret; do not commit or echo
export LINKPOOL_URL=...
export SIMPLYVC_URL=...

python3 tools/lpbench/lp_batch_probe.py \
  --rpc alchemy="$ALCHEMY_URL" \
  --rpc linkpool="$LINKPOOL_URL" \
  --rpc simplyvc="$SIMPLYVC_URL" \
  --blocks 300 \
  --batch-sizes 1,2,4,8,16,32,50 \
  --runs 3 \
  --address 0x... \
  --topic0 0x...
```

Use `--from-block` / `--to-block` to replay a specific window, for example the
2026-11-04 incident. By default the probe picks a range 10,000 blocks behind the
lowest head across all endpoints, so every endpoint has the blocks.

Arguments worth knowing:

| Flag | Meaning |
|---|---|
| `--rpc LABEL=URL` | Endpoint to test; repeat for multiple. URLs are never printed. |
| `--blocks N` | Number of consecutive blocks to walk (default 300). |
| `--batch-sizes` | Comma-separated chunk sizes (default `1,2,4,8,16,32,50`). |
| `--runs N` | Repetitions per batch size; median is reported (default 3). |
| `--latest-offset` | How far behind head to start (default 10,000). |
| `--address`, `--topic0` | Production log filter. Without these, the probe warns that it is querying all logs. |
| `--seed` | Shuffles batch-size order to spread provider cache effects. |

The script uses only the Python standard library and keeps HTTP connections
alive. It maps batch responses by JSON-RPC id, so out-of-order responses are
fine.

## 4. Experiment protocol

### 4.1 Endpoints

At minimum:

- `Alchemy_3` — current preferred endpoint, ~10 ms mean per call
- `LinkPool_0` — current fallback, ~200 ms mean per call
- one more provider (`SimplyVC_3/4`, or Chainstack after its keys are rotated)

Run from a pod in the same region/cluster as the VRF nodes. A laptop on VPN adds
WAN latency and compresses the relative differences.

### 4.2 Range and filter

- 300 consecutive blocks is enough for the A/B comparison and keeps the legacy
  baseline to ~2 minutes per endpoint.
- Use blocks at least 10,000 behind head so they are stable and every endpoint
  has them.
- Use the exact production VRF filter addresses/topics. Repeat with the BHS
  filter; the two nodes may register different filters.
- If the probe's legacy result is much faster than the production metric for the
  same endpoint, the filter or range is not representative. Calibrate first:
  LinkPool's legacy rate should be roughly 2–3 blocks/s.

### 4.3 Runs

- Run each batch size 3 times and take the median.
- Run the whole matrix more than once at different times of day if the endpoint
  is shared or noisy.
- Watch for 429s. If they appear, add a pause between runs or use a separate key;
  a throttled endpoint will make batching look worse than it is.

### 4.4 Safety

- The probe is read-only but consumes production RPC quota.
- Do not run the 2,000-block legacy baseline against an endpoint that is
  currently carrying live VRF traffic without coordinating with BCY. 300 blocks
  is usually harmless; 2,000 blocks is ~11 minutes of sustained calls.
- Do not paste RPC URLs into the PR, Slack, or shell history. Use environment
  variables or a secret.
- Start with small batch sizes and increase only after the previous size works.
  This also finds the provider's batch limit gradually instead of triggering it
  with a 50-call burst.

## 5. Interpreting results

Arbitrum's chain rate is ~3.7 blocks/s. Break-even is 3.7; target at least
10 blocks/s for steady-state headroom.

| LinkPool batch 8 result | Verdict |
|---|---|
| **>= 33 blocks/s** | L1 alone can recover a 2,000-block backlog within the 256-block BHS window. Chunk 8–16 is enough. |
| **10–33 blocks/s** | L1 fixes steady-state lag, but deep recovery still takes minutes. Rely on S1/S3/S4 for fast recovery; L1 is a safety net. |
| **4–10 blocks/s** | Marginal. Use chunk 16–32 or bounded concurrency, and prioritize S3/S4. |
| **< 4 blocks/s** | The provider is serializing batches. JSON-RPC batching will not fix the incident; use concurrency or drop L1's batching. |

Incident math, for context:

- At 2.3 blocks/s versus 3.7, lag grows 1.4 blocks/s.
- The 256-block BHS window closes in ~183 s (~3 minutes).
- S1's one-minute failback caps lag at ~84 blocks.
- A batch that reaches 6 blocks/s turns the deficit into a surplus and clears a
  256-block lag in ~40 s.

Deep catch-up, assuming `B(N)` stays roughly constant at 160 ms per batch:

| Batch | Throughput | 2,000-block backlog, net of chain growth |
|---:|---:|---:|
| 8 | ~25 blocks/s | ~94 s |
| 16 | ~50 blocks/s | ~43 s |
| 32 | ~100 blocks/s | ~21 s |
| 50 | ~156 blocks/s | ~13 s |

Those numbers are optimistic. The probe's actual `B(N)` is what matters. If the
provider's effective concurrency plateaus at 8, batch 50 gives no more
throughput than batch 8, and the extra latency only increases the reorg window.

## 6. Results so far

### 6.1 Zero-code baseline from existing metrics

From the same histogram family used in the RCA, at 2026-11-04 14:15 UTC (stable
LinkPool period). `CallContext` is a proxy for the header fetch; it includes
other calls, but on these VRF nodes it is dominated by
`eth_getBlockByNumber`.

| Endpoint | Mean header | Mean logs | Sequential capacity |
|---|---:|---:|---:|
| Alchemy_3 | 10.9 ms | 10.5 ms | ~46.6 blocks/s |
| LinkPool_0 | 212.8 ms | 197.3 ms | ~2.44 blocks/s |

The mean is the right predictor for a serial pipeline. The p50-based estimate
from the RCA (~3 blocks/s) is higher because the histogram has 50 ms buckets and
the tail is heavy.

PromQL:

```promql
sum by (nodeName, rpcCallName) (
  rate(evm_pool_rpc_node_rpc_call_time_sum{
    evmChainID="42161",
    service=~"cl-vrf-.*arb-mainnet.*",
    rpcCallName=~"CallContext|FilterLogs"
  }[10m])
)
/
sum by (nodeName, rpcCallName) (
  rate(evm_pool_rpc_node_rpc_call_time_count{
    evmChainID="42161",
    service=~"cl-vrf-.*arb-mainnet.*",
    rpcCallName=~"CallContext|FilterLogs"
  }[10m])
) / 1e6
```

### 6.2 Mock validation

20 ms per method call, 16 blocks:

| Provider behavior | Legacy | Batch 8 | Batch 16 |
|---|---:|---:|---:|
| Serial (`--batch-concurrency 1`) | ~20 blocks/s | ~20.8 (1.0x) | ~20.9 (1.0x) |
| Parallel (`--batch-concurrency 16`) | ~19.7 blocks/s | ~120 (6.1x) | ~199 (10.1x) |
| Max batch 4 | — | HTTP 413 | HTTP 413 |

This is the core finding: the same algorithm is a no-op or a 10x win depending
on the provider.

### 6.3 Public Arbitrum RPC control

`https://arb1.arbitrum.io/rpc`, 16 blocks ~10k behind head, unfiltered logs,
run from a laptop (absolute latency includes WAN):

| Mode | Blocks/s | Speedup |
|---|---:|---:|
| Legacy | 3.5 | 1.0x |
| Batch 1 | 3.8 | 1.1x |
| Batch 4 | 12.1 | 3.5x |
| Batch 8 | 23.1 | 6.6x |

A subsequent 64-block run returned HTTP 429. That is a useful reminder that
batching reduces round trips, not provider-side call accounting. Use a
production endpoint with enough quota and watch for 429s.

This public endpoint is a control, **not** a prediction for LinkPool. The
production endpoints still need to be measured.

## 7. What the probe does not measure

- Database insert time (`InsertLogsWithBlocks`) for large chunks
- head tracker / reorg handling
- the actual log poller code path, filters, or `skipEmptyBlocks` behavior
- provider rate limits under a sustained production-shaped load
- contention with the live VRF/BHS traffic on the same endpoint

Those require the full Go harness described in
[Recommended implementation shape](#recommended-implementation-shape), or a
canary node.

## L1 implementation review notes

The pseudocode in the RCA is directionally right, but it is not a patch. These
are the gaps to fix before implementation.

1. **`eth_getLogs` blockHash is singular.** Use a JSON-RPC batch of N independent
   calls, or a range query. A filter with N hashes is invalid.

2. **The current block's logs are queried every poll.** The existing loop calls
   `FilterLogs` for `currentBlock` before fetching `currentBlock+1`.
   `Test_getUnfinalizedLogs_includes_empty_finalized_when_skip_empty` expects
   N+1 log calls for N headers. The pseudocode starts at
   `currentBlock.Number+1`, silently skipping the anchor's logs. Preserve the
   anchor query or prove the DB always has them.

3. **`convertLogs` is not batch-safe.** It requires either one block or exactly
   one block per log. A chunk of 50 blocks with 7 logs triggers an
   `AssumptionViolation` and returns nothing. Convert per block or group logs by
   block hash.

4. **The replay DB check is missing.** `isReplay` compares each next block's
   hash with `SelectBlockByNumber` to avoid false reorgs. Batch mode needs a
   batched DB lookup or a per-header check before querying logs.

5. **`skipEmptyBlocks` breaks the 2000-block cap.** `maxUnfinalizedBlocks`
   counts saved blocks, not walked blocks. With empty blocks skipped, a batch
   loop can walk far more than 2000 blocks in one poll. Add a separate
   walked-block cap.

6. **Partial batch failures are normal.** `BatchCallContext` returns a nil
   top-level error even when individual `BatchElem.Error` fields are set. Check
   every element. If one log query fails, process the contiguous prefix and
   stop; do not skip the block or persist a gap.

7. **Endpoint-behind / null headers.** The head tracker may see block N via
   websocket while the HTTP endpoint is one or two blocks behind. A header batch
   can contain `null` at the tip. Treat null as "stop here and retry next poll",
   not as a fatal error.

8. **The reorg window widens.** Headers are fetched and verified, then logs are
   fetched. A reorg in between can make by-hash log queries return logs for
   orphaned blocks, or empty, depending on the provider. Keep chunks small (8–16)
   and decide explicitly how to treat "block not found" for a hash you just
   fetched. Add a mid-chunk reorg test.

9. **The 10-second timeout is closer than it looks.** `QueryTimeout` is 10 s, and
   `BatchCallContext` uses `largePayloadRPCTimeout` (also 10 s on Arbitrum). A
   serial batch of 50 at 160 ms is ~8 s, leaving almost no tail-latency headroom.
   A serial batch of 100 exceeds it. Chunk 8–16 is ~1.3–2.6 s serial.

10. **Metrics will undercount.** `BatchCallContext` calls `logResult` once for
    the whole batch, so `evm_pool_rpc_node_calls_total` increments by 1, not N,
    and the latency histogram records batch duration. The "endpoint in use"
    alert still works, but capacity/latency dashboards will be wrong. Add
    batch-size-aware metrics or per-element timing.

11. **Rate limits are per method call, not per HTTP request.** Batching reduces
    round trips, not provider compute. A 50-call burst can hit 429s even if the
    average rate is unchanged. Test sustained throughput, not one batch.

12. **Response size and memory.** A batch of 50 `eth_getLogs` responses can be
    large. Use the production filter, test a busy block window, and cap the
    response size.

Also: `fetchBlocks` already batches `eth_getBlockByNumber` for the backfill
path, but it appends a latest/finalized validation request and an optional
chain-reference request. Reuse `newBlockReq` / `validateBlockResponse`, but write
a simpler unfinalized-header variant. Do not reuse `RPCBatchSize` blindly; it is
tied to `GetBlocks`. Add a separate option/feature flag.

## Recommended implementation shape

If the probe shows batching helps, a conservative first implementation is:

- batch headers in chunks of 8–16; verify the parent chain
- query logs for those headers with **bounded concurrency** (4–8) using ordinary
  single-block `eth_getLogs` calls, rather than a JSON-RPC batch of logs
- preserve the anchor log query, replay DB check, per-block `convertLogs`,
  deferred checkpoint, and a separate walked-block cap
- add a config/feature flag so one node can canary it and roll back
- test reorg mid-chunk, partial batch failure, null header, endpoint-behind,
  `skipEmptyBlocks`, replay, max walked cap, and exact log equality
- then run the real log poller against a historical range with a real DB,
  because this probe does not measure DB insert time or head-tracker overhead

At 160 ms and concurrency 4, log queries alone give ~25 blocks/s; headers are
cheap. That is already >6x headroom and simpler to reason about than a
50-element JSON-RPC batch.

## References

- Incident RCA: Arbitrum VRF / BHS missed requests after RPC failover,
  2026-11-04.
- `pkg/logpoller/log_poller.go` — `getUnfinalizedLogs`, `convertLogs`,
  `fetchBlocks`, `maxUnfinalizedBlocks`.
- `pkg/client/rpc_client.go` — `BatchCallContext`, `logResult`.
- `pkg/client/compatibility_helper.go` — `toFilterArg` (single `blockHash`).
- `pkg/client/evm_client.go` — `QueryTimeout` (10 s).
- `pkg/logpoller/log_poller_internal_test.go` —
  `Test_getUnfinalizedLogs_includes_empty_finalized_when_skip_empty`.
