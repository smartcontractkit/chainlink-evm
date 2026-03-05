# OEV-987 — Does a node send a TX after receiving bids from Atlas?

**Status:** Investigation complete — answer is No, and here's exactly why.

---

## 1. The question

> "Check if we can carve out if NOPs are sending a tx when they receive a metacall with bids
> (i.e. more than 0 bids received from Atlas → then tx should be sent).
> We have metrics for bids and afterwards a TxMessage is sent."

Can we, by inspecting existing metrics / events, determine: when Atlas returned ≥1 bid,
did the node subsequently send a transaction?

---

## 2. Relevant code path (read before continuing)

All code lives in `chainlink-evm` (`/Users/ggerritsen/dev/cll/chainlink-evm`).

### 2a. The Atlas call and bid counting

`pkg/txm/clientwrappers/dualbroadcast/meta_client.go` — `SendRequest()`:

```
Atlas responds
  ├── error message containing "no solver operations…"
  │     → RecordBidsReceived(ctx, 0)
  │       return nil, nil          ← no error, but nil result
  ├── error message (other)
  │     → return nil, error        ← ErrAuction path
  ├── result == nil
  │     → return nil, nil
  └── result with SOS (solverOperations)
        → RecordBidsReceived(ctx, len(SOS))   ← bids > 0 recorded here
          → VerifyResponse(…)
```

`meta_client.go` L391–402:
```go
if strings.Contains(response.Error.ErrorMessage, NoSolverOps) || … {
    a.metrics.RecordBidsReceived(ctx, 0)
    return nil, nil
}
…
a.metrics.RecordBidsReceived(ctx, len(response.Result.SOS))
```

### 2b. What happens after bids > 0

Back in `SendTransaction()` (same file):

```go
meta, err := a.SendRequest(…)
if err != nil { … return ErrAuction }  // bids == 0 case handled inside SendRequest

if meta != nil {                        // ← bids > 0 path
    if err := a.SendOperation(…); err != nil {
        a.metrics.RecordSendOperationError(ctx)
        a.metrics.emitAtlasError(ctx, "send_operation", …)
        return err          // ← tx NOT sent; error bubbles up
    }
    return nil              // ← SUCCESS: tx was sent
}
// meta == nil means no bids
return ErrNoBids
```

### 2c. What the TXM loop does with errors

`pkg/txm/txm.go` — `sendTransactionWithError()`:

```go
txErr := t.client.SendTransaction(ctx, tx, attempt)   // calls MetaClient.SendTransaction
if txErr != nil && t.errorHandler != nil {
    if err = t.errorHandler.HandleError(ctx, tx, txErr, …); err != nil {
        return    // ← early return: IncrementNumBroadcastedTxs NOT called
    }
}
// …
t.Metrics.IncrementNumBroadcastedTxs(ctx)    // only reached on success
t.Metrics.EmitTxMessage(…)                  // only reached on success
```

`pkg/txm/clientwrappers/dualbroadcast/meta_error_handler.go` — `HandleError()`:

```go
if (errors.Is(txErr, ErrNoBids) || errors.Is(txErr, ErrAuction)) && tx.AttemptCount == 1 {
    txStore.MarkTxFatal(ctx, tx, tx.FromAddress)
    setNonce(tx.FromAddress, *tx.Nonce)
    return fmt.Errorf("transaction with txID: %d marked as fatal", tx.ID)
    // ← non-nil error → sendTransactionWithError returns early
}
```

---

## 3. Answer

### Case A — bids == 0 (no solver ops)

1. `RecordBidsReceived(ctx, 0)` is emitted (OTEL histogram `meta_bids_per_transaction`)
2. `SendRequest` returns `nil, nil`
3. `SendTransaction` returns `ErrNoBids`
4. `HandleError` marks tx as fatal, returns non-nil error
5. `sendTransactionWithError` returns early
6. **`IncrementNumBroadcastedTxs` and `EmitTxMessage` are NOT called**
7. No `TxMessage` Beholder event is emitted

### Case B — bids > 0, `SendOperation` succeeds

1. `RecordBidsReceived(ctx, N)` (N > 0) is emitted
2. `SendOperation` calls `c.SendTransaction(ctx, signedTx)` → tx goes to chain
3. `SendTransaction` returns `nil` (no error)
4. `sendTransactionWithError` reaches `IncrementNumBroadcastedTxs` and `EmitTxMessage`
5. **`TxMessage` Beholder event IS emitted**

### Case C — bids > 0, `SendOperation` fails

1. `RecordBidsReceived(ctx, N)` (N > 0) is emitted
2. `SendOperation` fails, `emitAtlasError` fires (Beholder `FastLaneAtlasError` event)
3. `SendTransaction` returns a non-`ErrNoBids`/`ErrAuction` wrapped error
4. `HandleError` falls through and returns the error unchanged
5. `sendTransactionWithError` has `txErr != nil`; `HandleError` returned non-nil → early return
6. **`TxMessage` Beholder event is NOT emitted**

---

## 4. Current observability coverage

| Signal | Where | What it tells us |
|---|---|---|
| `meta_bids_per_transaction` (OTEL histogram) | `meta_metrics.go` `RecordBidsReceived` | Distribution of bid counts per Atlas request |
| `txm_num_broadcasted_transactions` (OTEL counter + Prometheus) | `metrics.go` `IncrementNumBroadcastedTxs` | Successfully sent txs (only emitted after `SendOperation` succeeds) |
| `TxMessage` (Beholder proto event) | `metrics.go` `EmitTxMessage` | Per-tx detail; only on successful broadcast |
| `FastLaneAtlasError` (Beholder proto event) | `meta_metrics.go` `emitAtlasError` | Errors in `SendRequest` or `SendOperation` |
| `meta_errors` (OTEL counter) | `meta_metrics.go` `RecordSendRequestError` / `RecordSendOperationError` | Count of each error type |

**Correlation gap:** There is no single event or shared identifier that links a specific
`RecordBidsReceived` call to the `TxMessage` (or absence of one) that follows for the
same transaction. The bid histogram and the broadcast counter/event are independent.

---

## 5. Options for answering the question

The goal is: **every time bids > 0 is received from Atlas, we want to be able to verify
that a metacall was subsequently sent**. Ideally this is an alertable, Prometheus-queryable
signal.

---

### Option A — Two new OTEL counters in `meta_metrics.go` (recommended, no proto changes)

Add two counters to `MetaMetrics`:

- `meta_auctions_with_bids_total` — incremented immediately after `RecordBidsReceived`
  when `bidCount > 0`
- `meta_metacalls_sent_total` — incremented inside `SendOperation` after
  `c.SendTransaction` returns without error

Both labelled with `chainID` (and optionally `fromAddress` for per-NOP resolution).

**PromeQL alert:**
```promql
# Rate of auctions that got bids but metacall wasn't sent
increase(meta_auctions_with_bids_total[5m])
- increase(meta_metacalls_sent_total[5m])
```
Alert if this delta is consistently > 0.

**Why this works:** both counters are incremented inside `MetaClient`, which is the only
place that knows about both bids and metacall success. No plumbing through the TXM loop,
no proto changes, no release pipeline.

**Failure coverage:**
| Failure path | `meta_auctions_with_bids_total` | `meta_metacalls_sent_total` |
|---|---|---|
| `VerifyResponse` fails after bids counted | +1 | 0 |
| `VerifyMetadata` returns nil,nil (empty SOPs after decode) | +1 | 0 |
| `SendOperation` RPC fails | +1 | 0 |
| Everything succeeds | +1 | +1 |

**Files to change** (all in `chainlink-evm`):

| File | Change |
|---|---|
| `pkg/txm/clientwrappers/dualbroadcast/meta_metrics.go` | Add `auctionsWithBids` and `metacallsSent` counters; add `RecordAuctionWithBids()` and `RecordMetacallSent()` methods |
| `pkg/txm/clientwrappers/dualbroadcast/meta_client.go` | Call `RecordAuctionWithBids()` after `RecordBidsReceived` when count > 0; call `RecordMetacallSent()` after `c.SendTransaction` succeeds in `SendOperation` |
| `pkg/txm/clientwrappers/dualbroadcast/meta_metrics_test.go` | Add tests for the two new methods |

**Effort:** Small. Pure `chainlink-evm` change, no proto pipeline, no `chainlink` core bump needed.

---

### Option B — Add `bid_count` to the `FastLaneAtlasError` Beholder event

Field 8 is currently unused in `fastlane_atlas_error.proto`. Add `int32 bid_count = 8;`
and populate it when emitting an Atlas error after bids > 0 were received.

This makes every failure event self-describing: you can query Beholder for
`FastLaneAtlasError` events where `bid_count > 0` and know exactly how many bids were
lost per incident.

**Limitation:** only covers the failure path. Doesn't give you a counter for successful
metacalls with bids. You'd still need to compare against the success side separately.

**Effort:** Requires the proto → chainlink-evm release pipeline (see context doc §3).
The field slot (8) is free so no wire-format migration needed.

**Files to change:**

| Repo | File | Change |
|---|---|---|
| `chainlink-protos` | `svr/v1/fastlane_atlas_error.proto` | Add `int32 bid_count = 8;` |
| `chainlink-protos` | `svr/v1/fastlane_atlas_error.pb.go` | Regenerate (`task proto:gen:svr`) |
| `chainlink-evm` | `pkg/txm/clientwrappers/dualbroadcast/meta_metrics.go` | Add `bidCount int32` param to `emitAtlasError` |
| `chainlink-evm` | `pkg/txm/clientwrappers/dualbroadcast/meta_client.go` | Pass bid count to `emitAtlasError` calls |
| `chainlink-evm` | `pkg/txm/clientwrappers/dualbroadcast/meta_metrics_test.go` | Assert `bid_count` field in existing tests |

---

### Option C — Add `bid_count` to `TxMessage` Beholder event

Add `int32 bid_count = 9;` to `beholder_tx_message.proto` (field 8 is taken by
`dual_broadcast_params`). Propagate the count from `SendRequest` → `SendTransaction` →
`sendTransactionWithError` → `EmitTxMessage`.

This makes the success event self-describing: every `TxMessage` from an Atlas secondary
tx carries how many bids were in the winning metacall.

**Limitation:** same proto pipeline as Option B, but also requires plumbing the bid count
through `MetaClient.SendTransaction` return value → `txm.go` → `EmitTxMessage`. The TXM
core loop currently has no concept of bid count; adding it requires touching the
`sendTransactionWithError` signature or storing the count on the `Transaction`/`Attempt`
struct.

**Effort:** Higher than A or B.

---

### Option D — Structured logs only (no code changes)

`meta_client.go` logs at Info in `SendOperation`:
```go
a.lggr.Infow("Intercepted attempt for tx", "txID", tx.ID, "hash", signedTx.Hash(), …)
```
And `SendRequest` logs the full response JSON including solver operations. If logs are
shipped to a searchable store (Loki, etc.) this is queryable today with no code changes.

**Limitation:** log availability varies per NOP deployment; not alertable in Prometheus.

---

## 6. Recommendation

**Do Option A first.** It is self-contained in `chainlink-evm`, requires no proto
pipeline, and directly produces the Prometheus alert you want. Options B and C are
additive improvements for richer Beholder data and can follow in a subsequent PR.

The alert to build toward:
```promql
increase(meta_auctions_with_bids_total[5m]) - increase(meta_metacalls_sent_total[5m]) > 0
```

---

## 7. Open questions before implementing

1. Should `fromAddress` (the NOP's EOA) be a label on both new counters so the alert
   fires per-NOP? Or is `chainID` sufficient granularity?
2. Is there a naming convention for new OTEL metric names in `chainlink-evm`
   (prefix, snake_case, `_total` suffix)?

---

## 8. Implementation plan — Option A

**Decision:** Option A approved. `fromAddress` will be added as a label on both counters.

### Metric names and labels

Follows the existing convention in `meta_metrics.go` (`meta_` prefix, snake_case, no
`_total` suffix, label keys in camelCase):

| Metric | Type | Labels | Incremented when |
|---|---|---|---|
| `meta_auctions_with_bids` | `Int64Counter` | `chainID`, `fromAddress` | `RecordBidsReceived` is called with `bidCount > 0` |
| `meta_metacalls_sent` | `Int64Counter` | `chainID`, `fromAddress` | `c.SendTransaction` returns `nil` inside `SendOperation` |

### PromQL alert

```promql
increase(meta_auctions_with_bids[5m]) - increase(meta_metacalls_sent[5m]) > 0
```

### Exact changes

**`pkg/txm/clientwrappers/dualbroadcast/meta_metrics.go`**
- Add `auctionsWithBids metric.Int64Counter` and `metacallsSent metric.Int64Counter` fields to `MetaMetrics`
- Register both in `NewMetaMetrics`
- Add `RecordAuctionWithBids(ctx, fromAddress)` method
- Add `RecordMetacallSent(ctx, fromAddress)` method

**`pkg/txm/clientwrappers/dualbroadcast/meta_client.go`**
- In `SendRequest`: after `RecordBidsReceived(ctx, len(response.Result.SOS))`, if count > 0 call `RecordAuctionWithBids(ctx, tx.FromAddress.Hex())`
- In `SendOperation`: after `c.SendTransaction` returns nil, call `RecordMetacallSent(ctx, tx.FromAddress.Hex())`

**`pkg/txm/clientwrappers/dualbroadcast/meta_metrics_test.go`**
- Add tests for `RecordAuctionWithBids` and `RecordMetacallSent` (they don't panic, correct labels)

### Metrics and alert

| Metric | Type | Labels | Meaning |
|---|---|---|---|
| `meta_auctions_with_bids` | `Int64Counter` | `chainID`, `fromAddress` | Atlas returned ≥1 bid; a metacall should follow |
| `meta_metacalls_sent` | `Int64Counter` | `chainID`, `fromAddress` | Metacall was successfully submitted to the network |

**PromQL alert** — fires when a NOP receives bids but does not send a metacall:

```promql
increase(meta_auctions_with_bids[5m])
- increase(meta_metacalls_sent[5m])
> 0
```

Scope by `fromAddress` to identify the specific NOP, by `chainID` to identify the chain.

**What a gap means:** one of three failure paths occurred after bids were counted —
`VerifyResponse` rejected the metacall (security check failed), `VerifyMetadata` found
empty SOPs after decoding (Atlas returned inconsistent data), or `SendOperation`'s RPC
call to the network failed. In the latter two cases a `FastLaneAtlasError` Beholder event
is also emitted with the details.

### Status

- [x] Investigation complete
- [x] Options documented
- [x] Option A approved
- [x] Plan updated with metric names, labels, exact file changes
- [x] Implementation
- [x] Tests (all pass: `go test ./pkg/txm/clientwrappers/dualbroadcast/...`)
- [x] Metrics and alert documented
