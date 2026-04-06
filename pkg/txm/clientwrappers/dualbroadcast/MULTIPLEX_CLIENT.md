# Multiplex Client

The multiplex client sends each secondary (SVR/OEV) transaction to **two** OFA backends simultaneously: Flashbots MEV-Share (primary) and Nova RPC (secondary). This provides redundancy — if one backend underperforms or goes down, the other still has the transaction.

## How it works

```
         multiplexClient (implements txm.Client)
                │
        ┌───────┴───────┐
        │               │
  FlashbotsClient    novaClient
   (primary)         (secondary, fire-and-forget)
```

- **Primary (Flashbots)** determines the outcome for TXM. Its result is returned to the caller.
- **Secondary (Nova)** runs in a background goroutine. Errors are logged but never affect TXM's success/failure path.
- **Nonce queries** (`PendingNonceAt`, `NonceAt`) are routed exclusively to the primary.
- Each client manages its own HTTP timeout internally.

## Configuration

```toml
[EVM.Transactions.TransactionManagerV2]
Enabled = true
DualBroadcast = true
CustomURL = 'https://relay.flashbots.net'
CustomURLSecondary = 'https://eth.novarpc.xyz?api_key=YOUR_KEY'
Bundles = true
```

`CustomURLSecondary` must be a Nova RPC endpoint. The selector rejects non-Nova URLs for the secondary.

## Metrics

Both backends emit unified `ofa_*` metrics with a `backend` label (`"flashbots"` or `"nova"`):

| Metric | Type | Key labels |
|--------|------|------------|
| `ofa_send_tx_status` | counter | `backend`, `status` (`success`/`error`) |
| `ofa_send_tx_latency` | histogram (ms) | `backend`, `status` |
