# Beholder Metrics Assessment

**Date:** 2026-08-14
**Scope:** `pkg/logpoller`, `pkg/client`, `pkg/txm`, `pkg/txmgr`, `pkg/heads`
**Objective:** In the event of a production incident, how much Beholder metric data exists to determine "what's going on"?

---

## Executive summary

Beholder metrics in these packages are emitted through `beholder.GetMeter()`, an OpenTelemetry
meter that flows into the Beholder / Prometheus pipeline. Coverage is **very uneven across the
five services**:

| Service | Beholder gauges/counters/histograms | Attribute (label) richness | Diagnostic readiness |
|---|---|---|---|
| `pkg/client` (RPC pool) | 4 | **Good** — per-node, per-call-name, per-RPC-domain | Strong |
| `pkg/logpoller` | 1 | Minimal — chainID only | Limited |
| `pkg/txm` (TXMv2) | 7 meter + 2 histogram | **Weak — most emit zero labels** | Weak |
| `pkg/txmgr` (v1 TXM + framework) | 15 | Good — chainID + selectors | Strong |
| `pkg/heads` | **0** | n/a | **None** |

**Bottom line:** A production problem in **RPC/`pkg/client`** and **TXM/`pkg/txmgr`** is generally
diagnosable from metrics alone. A problem in **`pkg/logpoller`** has almost no metric signal (a
single gauge). A problem in **`pkg/txm`** is only partially diagnosable because the Beholder
exports are largely **unlabeled**. A problem in **`pkg/heads`** is **not diagnosable via metrics at
all** — there are no metrics in the package, only logs.

---

## Methodology

- Enumerated every `beholder.GetMeter().*` registration in each package (plus the
  `chainlink-framework/metrics` helpers that `pkg/client` and `pkg/txmgr` embed, since those also
  emit via `beholder.GetMeter()`).
- Recorded metric name, OTEL instrument type, attributes/labels attached at emission time, and the
  call frequency / trigger.
- Assessed *diagnostic strength*: how many independent axes (chain, node, call, stage, address,
  time) each metric separates data along, and whether the emitted series is sufficient to answer
  "what changed / where / when" for common production failures.

> Note: Many metrics are dual-emitted to **both** Beholder (OTEL) and legacy Prometheus
> (`promauto` registries). Where they disagree (notably in `pkg/txm`), this report calls it out,
> because the Beholder side is what an operator queries for unified OTEL-based observability.

---

## 1. `pkg/client` (RPC client / node pool)

`pkg/client/metrics.go` + `chainlink-framework/metrics` `RPCClientMetrics`.

### Inventory

| Beholder metric | OTEL type | Attributes |
|---|---|---|
| `evm_pool_rpc_node_calls_total` | Int64Counter | `chainFamily`, `chainID`, `nodeName`, `rpcDomain`, `callName` |
| `evm_pool_rpc_node_calls_success` | Int64Counter | same as above |
| `evm_pool_rpc_node_calls_failed` | Int64Counter | same as above |
| `rpc_call_latency` | Float64Histogram | `chainFamily`, `chainID`, `rpcDomain`, `isSendOnly`, `success`, `callName` |

Emitted per RPC call in `rpc_client.go:377-383` — i.e. **high-frequency, per-request**.

### Diagnostic strength — Strong

- **Per-node** resolution (`nodeName`, sanitized `rpcDomain`) lets you isolate a single bad RPC
  provider, a counterparty domain, or a single call type (`callName`, e.g. `eth_getLogs`).
- `success` / separate failed counter supports **error-rate by node and by call**.
- `rpc_call_latency` detects slow/unresponsive nodes.
- Together these answer the three canonical questions well:
  - *Where?* → `nodeName` / `rpcDomain`.
  - *What?* → `callName` (or `isSendOnly`).
  - *When / how bad?* → counters + latency histogram over time.

**This is the best-instrumented service in scope.**

### Notes

- `nodeName` and `rpcDomain` are **high-cardinality** labels; fine for Prometheus/OTEL as long as
  the node count stays bounded (it is — one per configured RPC URL). Watch out for unbounded
  `rpcDomain` growth if URLs are never pruned.
- If the RPC pool is misconfigured (bad URLs), these still emit useful failure signal because they
  are recorded regardless.

---

## 2. `pkg/logpoller`

`pkg/logpoller/metrics.go`.

### Inventory

| Beholder metric | OTEL type | Attributes |
|---|---|---|
| `evm_log_poller_last_processed_block` | Int64Gauge | `chainFamily`, `chainID` |

Emitted once per poll tick in `log_poller.go:760` (typically every poll interval / new block), and
**only when there is already a "last processed" block** — on the first-ever poll of a new chain the
gauge is never set.

### Diagnostic strength — Limited

A single gauge can only answer: *"Is the poller stuck?"* — compare
`evm_log_poller_last_processed_block` against the chain head. That is genuinely useful and is the
documented intent ("main purpose is to signal if the log poller is stuck").

But any other poller failure is invisible in metrics:

- **No error/success counters** for RPC calls, query failures, or save failures in `PollAndSaveLogs`.
- **No latency** on log queries (`eth_getLogs` is a classic slow-path and reorg source).
- **No per-filter/contract** resolution — a single contract silently missing logs is invisible
  (`chainID`/`chainFamily` only).
- **No reorgs / reorg depth** tracked.
- **No backfill progress / concurrency / backpressure.**

**For anything other than "poller is entirely stuck", you have to dig into logs.**

> Note: `chainlink-framework/metrics/logpoller.go` exists and offers additional Beholder logpoller
> metrics (query latency, reorgs, etc.), but it is **not wired in** in this packet's
> `pkg/logpoller` — only the single gauge is used.

---

## 3. `pkg/txm` (TXMv2)

`pkg/txm/metrics.go` + `pkg/txm/clientwrappers`.

### Inventory (core TXM metrics)

| Beholder metric | OTEL type | Attributes actually emitted | Prometheus label |
|---|---|---|---|
| `txm_num_broadcasted_transactions` | Int64Counter | **none** | `chainID` |
| `txm_num_confirmed_transactions` | Int64Counter | **none** | `chainID` |
| `txm_num_nonce_gaps` | Int64Counter | **none** | `chainID` |
| `txm_time_until_tx_confirmed` | Float64Histogram | **none** | `chainID` |
| `txm_reached_max_attempts` | Int64Gauge | **none** | `chainID` |
| `txm_rpc_nonce` | Int64Gauge | **none** | `chainID`, `address` |
| `txm_transaction_lifecycle_failure_total` | Int64Counter | `chainID`, `stage` | n/a |

Emitted from the TXM loop in `pkg/txm/txm.go` (broadcast, confirmation, nonce tracking, max-attempt
threshold) — i.e. per transaction lifecycle event.

### Diagnostic strength — Weak

The OTEL/Beholder side is **dramatically less useful than the dual-emitted Prometheus side**:

- Most `Add`/`Record` calls in `pkg/txm/metrics.go` (`IncrementNumBroadcastedTxs:165`,
  `IncrementNumConfirmedTxs:170`, `IncrementNumNonceGaps:175`, `ReachedMaxAttempts:184`,
  `RecordTimeUntilTxConfirmed:189`, `SetRPCNonce:194`) pass **no attributes at all**. In an OTEL
  world these collapse to a **single global series per metric** — no `chainID`, and critically
  `txm_rpc_nonce` loses its `address` dimension, so you cannot tell *which* sending address/keystore
  is stuck on a nonce.
- Only `txm_transaction_lifecycle_failure_total` carries a selector (`stage`) and `chainID` — but
  `stage` is a **fixed enum** (create, broadcast, nonce_at, …) with **no per-attempt error reason or
  per-address**, so two different root causes look identical.
- The Prometheus registries **do** label these correctly (`chainID`, and `address` for rpc_nonce), so
  the raw data is being captured — it just isn't surfaced through Beholder. This is a **plumbing /
  labeling bug**, not missing data collection.

What you *can* still answer from Beholder `txm` metrics: coarse *is the TXM broadcasting at all*,
*is it confirming*, *is it hitting max attempts*, and *lifecycle failure by stage*.

What you **cannot** easily answer: *which chain*, *which address*, *which attempt error caused the
failure*. For a stuck-keystore or per-chain TXM incident, look to the Prometheus side or logs.

### Side-car metrics in `pkg/txm/clientwrappers`

| Beholder metric | OTEL type | Attributes |
|---|---|---|
| `txm_multicall_duration_ms` | Float64Histogram | `chainID`, `method`, `blockTag`, `success`, `timedOut` |
| `ofa_send_tx_status` | Int64Counter | `chainID`, `backend`, `status` |
| `ofa_send_tx_latency` | Int64Histogram | `chainID`, `backend`, `status` |
| `meta_endpoint_status_codes` | Int64Counter | `chainID`, `statusCode`, `feedAddress` |
| `meta_endpoint_latency` | Int64Histogram | `chainID`, `feedAddress` |
| `meta_bids_per_transaction` | Int64Histogram | `chainID`, `feedAddress` |
| `meta_errors` | Int64Counter | `chainID`, `errorType`, `feedAddress` |

These are **well-labeled** (chainID + backend/statusCode/feedAddress) and are good diagnostics for
the OFA / Meta dual-broadcast paths. `chainID` here is a string label, not a real dimension for the
core `txm_*` metrics — a sign of inconsistent conventions.

---

## 4. `pkg/txmgr` (v1 TXM / legacy) + framework

`pkg/txmgr/metrics.go` + `chainlink-framework/metrics` `GenericTxmMetrics`.

### Inventory

| Beholder metric | OTEL type | Attributes |
|---|---|---|
| `tx_manager_num_successful_transactions` | Int64Counter | `chainID` |
| `tx_manager_num_tx_reverted` | Int64Counter | `chainID` |
| `tx_manager_fwd_tx_count` | Int64Counter | `chainID`, `successful` |
| `tx_manager_tx_attempt_count` | Float64Gauge | `chainID` |
| `tx_manager_num_finalized_transactions` | Int64Counter | `chainID` |
| `txm_pending_tx_queue_utilization` | Float64Gauge | `chainID` |
| `tx_manager_tx_oldest_non_terminal_age_seconds` | Float64Gauge | `chainID` |
| `tx_manager_num_broadcasted` (framework) | Int64Counter | `chainID` |
| `tx_manager_time_until_tx_broadcast` (framework) | Float64Histogram | `chainID` |
| `tx_manager_num_gas_bumps` (framework) | Int64Counter | `chainID` |
| `tx_manager_gas_bump_exceeds_limit` (framework) | Int64Counter | `chainID` |
| `tx_manager_num_confirmed_transactions` (framework) | Int64Counter | `chainID` |
| `tx_manager_time_until_tx_confirmed` (framework) | Float64Histogram | `chainID` |
| `tx_manager_blocks_until_tx_confirmed` (framework) | Float64Histogram | `chainID` |
| `tx_manager_attempt_error_total` (framework) | Int64Counter | `chainID`, `senderAddress`, `cause` |

Emitted from `pkg/txmgr/finalizer.go`, `evm_tx_store.go`, and the framework's own txm components at
transaction lifecycle milestones and on a poll loop (queue utilization, attempt count, oldest age).

### Diagnostic strength — Strong

- **Uniquely good red-flag metrics:** `tx_manager_gas_bump_exceeds_limit` (gas bumping is
  failing = serious), `txm_pending_tx_queue_utilization`, and
  `tx_manager_tx_oldest_non_terminal_age_seconds` (purpose-built for stuck-tx alerts).
- **Per-address and per-cause for failures:** `tx_manager_attempt_error_total` carries
  `senderAddress` and `cause` (with the `insufficient_funds` case already handled) — plugging the
  exact gap that `pkg/txm`'s lifecycle-failure metric has.
- Consistent `chainID` labeling across the whole set; the framework set is uniformly and correctly
  attributed.

The only weaknesses: `successful`/`cause` nuance is coarse on some counters (counts per confirmation
which can be inflated on re-orgs, noted in the metric Help strings), and there's no per-transaction
error message — but for *incident triage* this set is more than enough.

---

## 5. `pkg/heads` — **No metrics**

I searched `pkg/heads` for `beholder`, `GetMeter`, `metric`, `prometheus`, `promauto`, `otel`,
`gauge`, and `counter`. **Zero matches in production code** (the only hits are a test-only
`zap` observer and a comment). There is no `metrics.go` in `pkg/heads`.

`pkg/heads` (broadcaster, heads tracker, saver, ORM) is the thin consumer-visible edge of the head
/ block-subscription path. It participates in:

- emitting new-head events to every downstream consumer (log poller, TXM, jobs),
- tracking head latency / delivery,
- enforcing finality / reorg handling (`broadcaster.go`, `heads.go`, `tracker.go`).

With **zero** Beholder/Prometheus metrics, a production issue here (e.g. head broadcast stall, a
consumer holding the lock, delayed delivery, missed heads during a reorg) is **undetectable via
metrics**. Diagnosis would rely entirely on logs and downstream symptoms (e.g. the log poller gauge
stalling, or TXM nonce gaps), which is indirect and slow.

**This is the single largest observability gap in the scope.** Consider adding at minimum:
- a `heads` newest-block gauge (for lag vs. chain head),
- a head-delivery-latency histogram and a broadcast/consumer processing-time histogram,
- head-received and head-consumer-count counters,
- lag/clear-up metrics like logpoller's.

---

## Cross-cutting findings & recommendations

1. **Inconsistent attribute conventions.** `pkg/txm`/`pkg/txmgr` use per-`Add`/`Record`
   `metric.WithAttributes(...)`; `pkg/txm` also embeds a `metrics.Labeler` that is *never applied*,
   and passes `chainID` on some metrics but not others. Standardize on applying labels uniformly
   (ideally via the `Labeler`) so every series is at least `chainID`-dimensioned.

2. **`pkg/txm` Beholder exports are effectively unlabeled.** This is the highest-value, lowest-effort
   fix: add `chainID` (and `address` for `txm_rpc_nonce`) attributes to the six unlabeled TXMv2
   metrics (`metrics.go:165,170,175,184,189,194`). The dual-emitted Prometheus side proves the data
   is being captured — it just isn't reaching Beholder usefully.

3. **Log poller is a single-signal-only service.** Add query-latency, error counter, reorg, and
   per-contract dimensions (the framework's `logpoller.go` is already there but unconnected).

4. **`pkg/heads` has zero coverage** — most urgent greenfield addition.

5. **Cardinality guardrails:** `nodeName`/`rpcDomain`/`senderAddress`/`feedAddress`/`fromAddress`
   are the high-cardinality labels; they're bounded by configuration but should be monitored for
   unbounded growth.

### Recommended priority order

| Priority | Action | Impact |
|---|---|---|
| 1 | Add labels to `pkg/txm` Beholder metrics (`metrics.go`) | Turns weak TXM signal into strong |
| 2 | Add metrics to `pkg/heads` | Closes the only zero-coverage service |
| 3 | Wire framework logpoller metrics into `pkg/logpoller` | Adds reorg/latency/error signal |
| 4 | Standardize `Labeler`/attribute usage across all five packages | Keeps series consistently queryable |
