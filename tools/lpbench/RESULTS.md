# L1 probe results — 2026-11-06

Short summary of the first real-endpoint run of
[`lp_batch_probe.py`](lp_batch_probe.py) against Arbitrum mainnet VRF RPC
endpoints.

## How it was run

```bash
python3 tools/lpbench/lp_batch_probe.py \
  --rpc alchemy="$ALCHEMY_URL" \
  --rpc linkpool="$LINKPOOL_URL" \
  --blocks 300 \
  --batch-sizes 1,2,4,8,16,32,50 \
  --runs 3 \
  --address 0x3C0Ca683b403E37668AE3DC4FB62F4B29B6f7a3e \
  --topic0 0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e \
  --topic0 0xaeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b7
```

- 300 consecutive finalized Arbitrum blocks, ~10k behind head.
- Production VRF v2.5 coordinator filter
  (`RandomWordsRequested` + `RandomWordsFulfilled`), which is the filter
  registered by both the VRF v2plus listener and the BHS feeder.
- Run from a laptop, **not** from the VRF node's region/cluster.

## Results

| Endpoint | Legacy | Batch 8 | Batch 16 | Batch 50 |
|---|---:|---:|---:|---:|
| Alchemy | 12.2 blocks/s (~41 ms/call) | 71.5 (5.9x) | 143.1 (11.8x) | 356.5 (29.3x) |
| LinkPool | 11.3 blocks/s (~44 ms/call) | 78.7 (7.0x) | 161.3 (14.3x) | 414.4 (36.8x) |

- All runs returned identical logs; no 413/429.
- Scaling is near-linear through batch 16. Effective concurrency is ~29–37 at
  batch 50, so the provider caps around there.
- Batch 16 is already ~39–43x Arbitrum's ~3.7 blocks/s.

## Caveat

The ~41–44 ms/call legacy latency is much lower than the incident's
~200 ms/call. This run was from a laptop, not the VRF node's region. Current
production metrics show Alchemy_3 at ~10.9 ms, Chainstack_0 at ~31.5 ms, and
4–5 blocks of lag; LinkPool is not carrying primary traffic. The absolute legacy
rate is therefore not production-faithful, but the **batch speedup** is the
transferable signal.

## Recommendation

Use chunk **8–16**, not 50. Batch 8 gives ~19–21x headroom; batch 16 gives
~39–43x and is better for catch-up. Batch 50 adds reorg-window, response-size,
and tail-latency risk for no operational benefit.

Next step: repeat from the VRF node's region/cluster to confirm the absolute
legacy latency before implementing L1.
