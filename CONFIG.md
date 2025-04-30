[//]: # (Documentation generated from docs.toml - DO NOT EDIT.)
This document describes the TOML format for configuration.
## Example

```toml
ChainID = '1' # Required

[[Nodes]]
Name = 'fake' # Required
WSURL = 'wss://foo.bar/ws'
HTTPURL = 'https://foo.bar' # Required

```

## Global
```toml
ChainID = '1' # Example
Enabled = true # Default
AutoCreateKey = true # Default
BlockBackfillDepth = 10 # Default
BlockBackfillSkip = false # Default
ChainType = 'arbitrum' # Example
FinalityDepth = 50 # Default
FinalityTagEnabled = false # Default
FlagsContractAddress = '0xae4E781a6218A8031764928E88d457937A954fC3' # Example
LinkContractAddress = '0x538aAaB4ea120b2bC2fe5D296852D948F07D849e' # Example
LogBackfillBatchSize = 1000 # Default
LogPollInterval = '15s' # Default
LogKeepBlocksDepth = 100000 # Default
LogPrunePageSize = 0 # Default
BackupLogPollerBlockDelay = 100 # Default
MinContractPayment = '10000000000000 juels' # Default
MinIncomingConfirmations = 3 # Default
NonceAutoSync = true # Default
NoNewHeadsThreshold = '3m' # Default
OperatorFactoryAddress = '0xa5B85635Be42F21f94F28034B7DA440EeFF0F418' # Example
RPCDefaultBatchSize = 250 # Default
RPCBlockQueryDelay = 1 # Default
FinalizedBlockOffset = 0 # Default
LogBroadcasterEnabled = true # Default
NoNewFinalizedHeadsThreshold = '0' # Default
```


### ChainID
```toml
ChainID = '1' # Example
```
ChainID is the EVM chain ID. Mandatory.

### Enabled
```toml
Enabled = true # Default
```
Enabled enables this chain.

### AutoCreateKey
```toml
AutoCreateKey = true # Default
```
AutoCreateKey, if set to true, will ensure that there is always at least one transmit key for the given chain.

### BlockBackfillDepth
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
BlockBackfillDepth = 10 # Default
```
BlockBackfillDepth specifies the number of blocks before the current HEAD that the log broadcaster will try to re-consume logs from.

### BlockBackfillSkip
```toml
BlockBackfillSkip = false # Default
```
BlockBackfillSkip enables skipping of very long backfills.

### ChainType
```toml
ChainType = 'arbitrum' # Example
```
ChainType is automatically detected from chain ID. Set this to force a certain chain type regardless of chain ID.
Available types: `arbitrum`, `celo`, `gnosis`, `hedera`, `kroma`, `metis`, `optimismBedrock`, `scroll`, `wemix`, `xlayer`, `zksync`

### FinalityDepth
```toml
FinalityDepth = 50 # Default
```
FinalityDepth is the number of blocks after which an ethereum transaction is considered "final". Note that the default is automatically set based on chain ID, so it should not be necessary to change this under normal operation.
BlocksConsideredFinal determines how deeply we look back to ensure that transactions are confirmed onto the longest chain
There is not a large performance penalty to setting this relatively high (on the order of hundreds)
It is practically limited by the number of heads we store in the database and should be less than this with a comfortable margin.
If a transaction is mined in a block more than this many blocks ago, and is reorged out, we will NOT retransmit this transaction and undefined behaviour can occur including gaps in the nonce sequence that require manual intervention to fix.
Therefore, this number represents a number of blocks we consider large enough that no re-org this deep will ever feasibly happen.

Special cases:
`FinalityDepth`=0 would imply that transactions can be final even before they were mined into a block. This is not supported.
`FinalityDepth`=1 implies that transactions are final after we see them in one block.

Examples:

Transaction sending:
A transaction is sent at block height 42

`FinalityDepth` is set to 5
A re-org occurs at height 44 starting at block 41, transaction is marked for rebroadcast
A re-org occurs at height 46 starting at block 41, transaction is marked for rebroadcast
A re-org occurs at height 47 starting at block 41, transaction is NOT marked for rebroadcast

### FinalityTagEnabled
```toml
FinalityTagEnabled = false # Default
```
FinalityTagEnabled means that the chain supports the finalized block tag when querying for a block. If FinalityTagEnabled is set to true for a chain, then FinalityDepth field is ignored.
Finality for a block is solely defined by the finality related tags provided by the chain's RPC API. This is a placeholder and hasn't been implemented yet.

### FlagsContractAddress
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
FlagsContractAddress = '0xae4E781a6218A8031764928E88d457937A954fC3' # Example
```
FlagsContractAddress can optionally point to a [Flags contract](../contracts/src/v0.8/Flags.sol). If set, the node will lookup that contract for each job that supports flags contracts (currently OCR and FM jobs are supported). If the job's contractAddress is set as hibernating in the FlagsContractAddress address, it overrides the standard update parameters (such as heartbeat/threshold).

### LinkContractAddress
```toml
LinkContractAddress = '0x538aAaB4ea120b2bC2fe5D296852D948F07D849e' # Example
```
LinkContractAddress is the canonical ERC-677 LINK token contract address on the given chain. Note that this is usually autodetected from chain ID.

### LogBackfillBatchSize
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
LogBackfillBatchSize = 1000 # Default
```
LogBackfillBatchSize sets the batch size for calling FilterLogs when we backfill missing logs.

### LogPollInterval
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
LogPollInterval = '15s' # Default
```
LogPollInterval works in conjunction with Feature.LogPoller. Controls how frequently the log poller polls for logs. Defaults to the block production rate.

### LogKeepBlocksDepth
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
LogKeepBlocksDepth = 100000 # Default
```
LogKeepBlocksDepth works in conjunction with Feature.LogPoller. Controls how many blocks the poller will keep, must be greater than FinalityDepth+1.

### LogPrunePageSize
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
LogPrunePageSize = 0 # Default
```
LogPrunePageSize defines size of the page for pruning logs. Controls how many logs/blocks (at most) are deleted in a single prune tick. Default value 0 means no paging, delete everything at once.

### BackupLogPollerBlockDelay
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
BackupLogPollerBlockDelay = 100 # Default
```
BackupLogPollerBlockDelay works in conjunction with Feature.LogPoller. Controls the block delay of Backup LogPoller, affecting how far behind the latest finalized block it starts and how often it runs.
BackupLogPollerDelay=0 will disable Backup LogPoller (_not recommended for production environment_).

### MinContractPayment
```toml
MinContractPayment = '10000000000000 juels' # Default
```
MinContractPayment is the minimum payment in LINK required to execute a direct request job. This can be overridden on a per-job basis.

### MinIncomingConfirmations
```toml
MinIncomingConfirmations = 3 # Default
```
MinIncomingConfirmations is the minimum required confirmations before a log event will be consumed.

### NonceAutoSync
```toml
NonceAutoSync = true # Default
```
NonceAutoSync enables automatic nonce syncing on startup. Chainlink nodes will automatically try to sync its local nonce with the remote chain on startup and fast forward if necessary. This is almost always safe but can be disabled in exceptional cases by setting this value to false.

### NoNewHeadsThreshold
```toml
NoNewHeadsThreshold = '3m' # Default
```
NoNewHeadsThreshold controls how long to wait after receiving no new heads before `NodePool` marks rpc endpoints as
out-of-sync, and `HeadTracker` logs warnings.

Set to zero to disable out-of-sync checking.

### OperatorFactoryAddress
```toml
OperatorFactoryAddress = '0xa5B85635Be42F21f94F28034B7DA440EeFF0F418' # Example
```
OperatorFactoryAddress is the address of the canonical operator forwarder contract on the given chain. Note that this is usually autodetected from chain ID.

### RPCDefaultBatchSize
```toml
RPCDefaultBatchSize = 250 # Default
```
RPCDefaultBatchSize is the default batch size for batched RPC calls.

### RPCBlockQueryDelay
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
RPCBlockQueryDelay = 1 # Default
```
RPCBlockQueryDelay controls the number of blocks to trail behind head in the block history estimator and balance monitor.
For example, if this is set to 3, and we receive block 10, block history estimator will fetch block 7.

CAUTION: You might be tempted to set this to 0 to use the latest possible
block, but it is possible to receive a head BEFORE that block is actually
available from the connected node via RPC, due to race conditions in the code of the remote ETH node. In this case you will get false
"zero" blocks that are missing transactions.

### FinalizedBlockOffset
```toml
FinalizedBlockOffset = 0 # Default
```
FinalizedBlockOffset defines the number of blocks by which the latest finalized block will be shifted/delayed.
For example, suppose RPC returns block 100 as the latest finalized. In that case, the CL Node will treat block `100 - FinalizedBlockOffset` as the latest finalized block and `latest - FinalityDepth - FinalizedBlockOffset` in case of `FinalityTagEnabled = false.`
With `EnforceRepeatableRead = true,` RPC is considered healthy only if its most recent finalized block is larger or equal to the highest finalized block observed by the CL Node minus `FinalizedBlockOffset.`
Higher values of `FinalizedBlockOffset` with `EnforceRepeatableRead = true` reduce the number of false `FinalizedBlockOutOfSync` declarations on healthy RPCs that are slightly lagging behind due to network delays.
This may increase the number of healthy RPCs and reduce the probability that the CL Node will not have any healthy alternatives to the active RPC.
CAUTION: Setting this to values higher than 0 may delay transaction creation in products (e.g., CCIP, Automation) that base their decision on finalized on-chain events.
PoS chains with `FinalityTagEnabled=true` and batched (epochs) blocks finalization (e.g., Ethereum Mainnet) must be treated with special care as a minor increase in the `FinalizedBlockOffset` may lead to significant delays.
For example, let's say that `FinalizedBlockOffset = 1` and blocks are finalized in batches of 32.
The latest finalized block on chain is 64, so block 63 is the latest finalized for CL Node.
Block 64 will be treated as finalized by CL Node only when chain's latest finalized block is 65. As chain finalizes blocks in batches of 32,
CL Node has to wait for a whole new batch to be finalized to treat block 64 as finalized.

### LogBroadcasterEnabled
```toml
LogBroadcasterEnabled = true # Default
```
LogBroadcasterEnabled is a feature flag for LogBroadcaster, by default it's true.

### NoNewFinalizedHeadsThreshold
```toml
NoNewFinalizedHeadsThreshold = '0' # Default
```
NoNewFinalizedHeadsThreshold controls how long to wait for new finalized block before `NodePool` marks rpc endpoints as
out-of-sync. Only applicable if `FinalityTagEnabled=true`

Set to zero to disable.

## Transactions
```toml
[Transactions]
Enabled = true # Default
ForwardersEnabled = false # Default
MaxInFlight = 16 # Default
MaxQueued = 250 # Default
ReaperInterval = '1h' # Default
ReaperThreshold = '168h' # Default
ResendAfterThreshold = '1m' # Default
```


### Enabled
```toml
Enabled = true # Default
```
Enabled is a feature flag for the Transaction Manager. This flag also enables or disables the gas estimator since it is dependent on the TXM to start it.

### ForwardersEnabled
```toml
ForwardersEnabled = false # Default
```
ForwardersEnabled enables or disables sending transactions through forwarder contracts.

### MaxInFlight
```toml
MaxInFlight = 16 # Default
```
MaxInFlight controls how many transactions are allowed to be "in-flight" i.e. broadcast but unconfirmed at any one time. You can consider this a form of transaction throttling.

The default is set conservatively at 16 because this is a pessimistic minimum that both geth and parity will hold without evicting local transactions. If your node is falling behind and you need higher throughput, you can increase this setting, but you MUST make sure that your ETH node is configured properly otherwise you can get nonce gapped and your node will get stuck.

0 value disables the limit. Use with caution.

### MaxQueued
```toml
MaxQueued = 250 # Default
```
MaxQueued is the maximum number of unbroadcast transactions per key that are allowed to be enqueued before jobs will start failing and rejecting send of any further transactions. This represents a sanity limit and generally indicates a problem with your ETH node (transactions are not getting mined).

Do NOT blindly increase this value thinking it will fix things if you start hitting this limit because transactions are not getting mined, you will instead only make things worse.

In deployments with very high burst rates, or on chains with large re-orgs, you _may_ consider increasing this.

0 value disables any limit on queue size. Use with caution.

### ReaperInterval
```toml
ReaperInterval = '1h' # Default
```
ReaperInterval controls how often the EthTx reaper will run.

### ReaperThreshold
```toml
ReaperThreshold = '168h' # Default
```
ReaperThreshold indicates how old an EthTx ought to be before it can be reaped.

### ResendAfterThreshold
```toml
ResendAfterThreshold = '1m' # Default
```
ResendAfterThreshold controls how long to wait before re-broadcasting a transaction that has not yet been confirmed.

## Transactions.AutoPurge
```toml
[Transactions.AutoPurge]
Enabled = false # Default
DetectionApiUrl = 'https://example.api.io' # Example
Threshold = 5 # Example
MinAttempts = 3 # Example
```


### Enabled
```toml
Enabled = false # Default
```
Enabled enables or disables automatically purging transactions that have been idenitified as terminally stuck (will never be included on-chain). This feature is only expected to be used by ZK chains.

### DetectionApiUrl
```toml
DetectionApiUrl = 'https://example.api.io' # Example
```
DetectionApiUrl configures the base url of a custom endpoint used to identify terminally stuck transactions.

### Threshold
```toml
Threshold = 5 # Example
```
Threshold configures the number of blocks a transaction has to remain unconfirmed before it is evaluated for being terminally stuck. This threshold is only applied if there is no custom API to identify stuck transactions provided by the chain.

### MinAttempts
```toml
MinAttempts = 3 # Example
```
MinAttempts configures the minimum number of broadcasted attempts a transaction has to have before it is evaluated further for being terminally stuck. This threshold is only applied if there is no custom API to identify stuck transactions provided by the chain. Ensure the gas estimator configs take more bump attempts before reaching the configured max gas price.

## Transactions.TransactionManagerV2
```toml
[Transactions.TransactionManagerV2]
Enabled = false # Default
BlockTime = '10s' # Example
CustomURL = 'https://example.api.io' # Example
DualBroadcast = false # Example
```


### Enabled
```toml
Enabled = false # Default
```
Enabled enables TransactionManagerV2.

### BlockTime
```toml
BlockTime = '10s' # Example
```
BlockTime controls the frequency of the backfill loop of TransactionManagerV2.

### CustomURL
```toml
CustomURL = 'https://example.api.io' # Example
```
CustomURL configures the base url of a custom endpoint used by the ChainDualBroadcast chain type.

### DualBroadcast
```toml
DualBroadcast = false # Example
```
DualBroadcast enables DualBroadcast functionality.

## BalanceMonitor
```toml
[BalanceMonitor]
Enabled = true # Default
```


### Enabled
```toml
Enabled = true # Default
```
Enabled balance monitoring for all keys.

## GasEstimator
```toml
[GasEstimator]
Mode = 'BlockHistory' # Default
PriceDefault = '20 gwei' # Default
PriceMax = '115792089237316195423570985008687907853269984665.640564039457584007913129639935 tether' # Default
PriceMin = '1 gwei' # Default
LimitDefault = 500_000 # Default
LimitMax = 500_000 # Default
LimitMultiplier = '1.0' # Default
LimitTransfer = 21_000 # Default
EstimateLimit = false # Default
BumpMin = '5 gwei' # Default
BumpPercent = 20 # Default
BumpThreshold = 3 # Default
BumpTxDepth = 16 # Example
EIP1559DynamicFees = false # Default
FeeCapDefault = '100 gwei' # Default
TipCapDefault = '1 wei' # Default
TipCapMin = '1 wei' # Default
```


### Mode
```toml
Mode = 'BlockHistory' # Default
```
Mode controls what type of gas estimator is used.

- `FixedPrice` uses static configured values for gas price (can be set via API call).
- `BlockHistory` dynamically adjusts default gas price based on heuristics from mined blocks.
- `L2Suggested` mode is deprecated and replaced with `SuggestedPrice`.
- `SuggestedPrice` is a mode which uses the gas price suggested by the rpc endpoint via `eth_gasPrice`.
- `Arbitrum` is a special mode only for use with Arbitrum blockchains. It uses the suggested gas price (up to `ETH_MAX_GAS_PRICE_WEI`, with `1000 gwei` default) as well as an estimated gas limit (up to `ETH_GAS_LIMIT_MAX`, with `1,000,000,000` default).

Chainlink nodes decide what gas price to use using an `Estimator`. It ships with several simple and battle-hardened built-in estimators that should work well for almost all use-cases. Note that estimators will change their behaviour slightly depending on if you are in EIP-1559 mode or not.

You can also use your own estimator for gas price by selecting the `FixedPrice` estimator and using the exposed API to set the price.

An important point to note is that the Chainlink node does _not_ ship with built-in support for go-ethereum's `estimateGas` call. This is for several reasons, including security and reliability. We have found empirically that it is not generally safe to rely on the remote ETH node's idea of what gas price should be.

### PriceDefault
```toml
PriceDefault = '20 gwei' # Default
```
PriceDefault is the default gas price to use when submitting transactions to the blockchain. Will be overridden by the built-in `BlockHistoryEstimator` if enabled, and might be increased if gas bumping is enabled.

(Only applies to legacy transactions)

Can be used with the `chainlink setgasprice` to be updated while the node is still running.

### PriceMax
```toml
PriceMax = '115792089237316195423570985008687907853269984665.640564039457584007913129639935 tether' # Default
```
PriceMax is the maximum gas price. Chainlink nodes will never pay more than this for a transaction.
This applies to both legacy and EIP1559 transactions.
Note that it is impossible to disable the maximum limit. Setting this value to zero will prevent paying anything for any transaction (which can be useful in some rare cases).
Most chains by default have the maximum set to 2**256-1 Wei which is the maximum allowed gas price on EVM-compatible chains, and is so large it may as well be unlimited.

### PriceMin
```toml
PriceMin = '1 gwei' # Default
```
PriceMin is the minimum gas price. Chainlink nodes will never pay less than this for a transaction.

(Only applies to legacy transactions)

It is possible to force the Chainlink node to use a fixed gas price by setting a combination of these, e.g.

```toml
EIP1559DynamicFees = false
PriceMax = 100
PriceMin = 100
PriceDefault = 100
BumpThreshold = 0
Mode = 'FixedPrice'
```

### LimitDefault
```toml
LimitDefault = 500_000 # Default
```
LimitDefault sets default gas limit for outgoing transactions. This should not need to be changed in most cases.
Some job types, such as Keeper jobs, might set their own gas limit unrelated to this value.

### LimitMax
```toml
LimitMax = 500_000 # Default
```
LimitMax sets a maximum for _estimated_ gas limits. This currently only applies to `Arbitrum` `GasEstimatorMode`.

### LimitMultiplier
```toml
LimitMultiplier = '1.0' # Default
```
LimitMultiplier is the factor by which a transaction's GasLimit is multiplied before transmission. So if the value is 1.1, and the GasLimit for a transaction is 10, 10% will be added before transmission.

This factor is always applied, so includes L2 transactions which uses a default gas limit of 1 and is also applied to `LimitDefault`.

### LimitTransfer
```toml
LimitTransfer = 21_000 # Default
```
LimitTransfer is the gas limit used for an ordinary ETH transfer.

### EstimateLimit
```toml
EstimateLimit = false # Default
```
EstimateLimit enables estimating gas limits for transactions. This feature respects the gas limit provided during transaction creation as an upper bound.

### BumpMin
```toml
BumpMin = '5 gwei' # Default
```
BumpMin is the minimum fixed amount of wei by which gas is bumped on each transaction attempt.

### BumpPercent
```toml
BumpPercent = 20 # Default
```
BumpPercent is the percentage by which to bump gas on a transaction that has exceeded `BumpThreshold`. The larger of `BumpPercent` and `BumpMin` is taken for gas bumps.

The `SuggestedPriceEstimator` adds the larger of `BumpPercent` and `BumpMin` on top of the price provided by the RPC when bumping a transaction's gas.

### BumpThreshold
```toml
BumpThreshold = 3 # Default
```
BumpThreshold is the number of blocks to wait for a transaction stuck in the mempool before automatically bumping the gas price. Set to 0 to disable gas bumping completely.

### BumpTxDepth
```toml
BumpTxDepth = 16 # Example
```
BumpTxDepth is the number of transactions to gas bump starting from oldest. Set to 0 for no limit (i.e. bump all). Can not be greater than EVM.Transactions.MaxInFlight. If not set, defaults to EVM.Transactions.MaxInFlight.

### EIP1559DynamicFees
```toml
EIP1559DynamicFees = false # Default
```
EIP1559DynamicFees torces EIP-1559 transaction mode. Enabling EIP-1559 mode can help reduce gas costs on chains that support it. This is supported only on official Ethereum mainnet and testnets. It is not recommended to enable this setting on Polygon because the EIP-1559 fee market appears to be broken on all Polygon chains and EIP-1559 transactions are less likely to be included than legacy transactions.

#### Technical details

Chainlink nodes include experimental support for submitting transactions using type 0x2 (EIP-1559) envelope.

EIP-1559 mode is enabled by default on the Ethereum Mainnet, but can be enabled on a per-chain basis or globally.

This might help to save gas on spikes. Chainlink nodes should react faster on the upleg and avoid overpaying on the downleg. It might also be possible to set `EVM.GasEstimator.BlockHistory.BatchSize` to a smaller value such as 12 or even 6 because tip cap should be a more consistent indicator of inclusion time than total gas price. This would make Chainlink nodes more responsive and should reduce response time variance. Some experimentation is required to find optimum settings.

Set with caution, if you set this on a chain that does not actually support EIP-1559 your node will be broken.

In EIP-1559 mode, the total price for the transaction is the minimum of base fee + tip cap and fee cap. More information can be found on the [official EIP](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1559.md).

Chainlink's implementation of EIP-1559 works as follows:

If you are using FixedPriceEstimator:
- With gas bumping disabled, it will submit all transactions with `feecap=PriceMax` and `tipcap=GasTipCapDefault`
- With gas bumping enabled, it will submit all transactions initially with `feecap=GasFeeCapDefault` and `tipcap=GasTipCapDefault`.

If you are using BlockHistoryEstimator (default for most chains):
- With gas bumping disabled, it will submit all transactions with `feecap=PriceMax` and `tipcap=<calculated using past blocks>`
- With gas bumping enabled (default for most chains) it will submit all transactions initially with `feecap = ( current block base fee * (1.125 ^ N) + tipcap )` where N is configurable by setting `EVM.GasEstimator.BlockHistory.EIP1559FeeCapBufferBlocks` but defaults to `gas bump threshold+1` and `tipcap=<calculated using past blocks>`

Bumping works as follows:

- Increase tipcap by `max(tipcap * (1 + BumpPercent), tipcap + BumpMin)`
- Increase feecap by `max(feecap * (1 + BumpPercent), feecap + BumpMin)`

A quick note on terminology - Chainlink nodes use the same terms used internally by go-ethereum source code to describe various prices. This is not the same as the externally used terms. For reference:

- Base Fee Per Gas = BaseFeePerGas
- Max Fee Per Gas = FeeCap
- Max Priority Fee Per Gas = TipCap

In EIP-1559 mode, the following changes occur to how configuration works:

- All new transactions will be sent as type 0x2 transactions specifying a TipCap and FeeCap. Be aware that existing pending legacy transactions will continue to be gas bumped in legacy mode.
- `BlockHistoryEstimator` will apply its calculations (gas percentile etc) to the TipCap and this value will be used for new transactions (GasPrice will be ignored)
- `FixedPriceEstimator` will use `GasTipCapDefault` instead of `GasPriceDefault` for the tip cap
- `FixedPriceEstimator` will use `GasFeeCapDefault` instaed of `GasPriceDefault` for the fee cap
- `PriceMin` is ignored for new transactions and `GasTipCapMinimum` is used instead (default 0)
- `PriceMax` still represents that absolute upper limit that Chainlink will ever spend (total) on a single tx
- `Keeper.GasTipCapBufferPercent` is ignored in EIP-1559 mode and `Keeper.GasTipCapBufferPercent` is used instead

### FeeCapDefault
```toml
FeeCapDefault = '100 gwei' # Default
```
FeeCapDefault controls the fixed initial fee cap, if EIP1559 mode is enabled and `FixedPrice` gas estimator is used.

### TipCapDefault
```toml
TipCapDefault = '1 wei' # Default
```
TipCapDefault is the default gas tip to use when submitting transactions to the blockchain. Will be overridden by the built-in `BlockHistoryEstimator` if enabled, and might be increased if gas bumping is enabled.

(Only applies to EIP-1559 transactions)

### TipCapMin
```toml
TipCapMin = '1 wei' # Default
```
TipCapMinimum is the minimum gas tip to use when submitting transactions to the blockchain.

(Only applies to EIP-1559 transactions)

## GasEstimator.DAOracle
```toml
[GasEstimator.DAOracle]
OracleType = 'opstack' # Example
OracleAddress = '0x420000000000000000000000000000000000000F' # Example
CustomGasPriceCalldata = '' # Default
```


### OracleType
```toml
OracleType = 'opstack' # Example
```
OracleType refers to the oracle family this config belongs to. Currently the available oracle types are: 'opstack', 'arbitrum', 'zksync', and 'custom_calldata'.

### OracleAddress
```toml
OracleAddress = '0x420000000000000000000000000000000000000F' # Example
```
OracleAddress is the address of the oracle contract.

### CustomGasPriceCalldata
```toml
CustomGasPriceCalldata = '' # Default
```
CustomGasPriceCalldata is optional and can be set to call a custom gas price function at the given OracleAddress.

## GasEstimator.LimitJobType
```toml
[GasEstimator.LimitJobType]
OCR = 100_000 # Example
OCR2 = 100_000 # Example
DR = 100_000 # Example
VRF = 100_000 # Example
FM = 100_000 # Example
Keeper = 100_000 # Example
```


### OCR
```toml
OCR = 100_000 # Example
```
OCR overrides LimitDefault for OCR jobs.

### OCR2
```toml
OCR2 = 100_000 # Example
```
OCR2 overrides LimitDefault for OCR2 jobs.

### DR
```toml
DR = 100_000 # Example
```
DR overrides LimitDefault for Direct Request jobs.

### VRF
```toml
VRF = 100_000 # Example
```
VRF overrides LimitDefault for VRF jobs.

### FM
```toml
FM = 100_000 # Example
```
FM overrides LimitDefault for Flux Monitor jobs.

### Keeper
```toml
Keeper = 100_000 # Example
```
Keeper overrides LimitDefault for Keeper jobs.

## GasEstimator.BlockHistory
```toml
[GasEstimator.BlockHistory]
BatchSize = 25 # Default
BlockHistorySize = 8 # Default
CheckInclusionBlocks = 12 # Default
CheckInclusionPercentile = 90 # Default
EIP1559FeeCapBufferBlocks = 13 # Example
TransactionPercentile = 60 # Default
```
These settings allow you to configure how your node calculates gas prices when using the block history estimator.
In most cases, leaving these values at their defaults should give good results.

### BatchSize
```toml
BatchSize = 25 # Default
```
BatchSize sets the maximum number of blocks to fetch in one batch in the block history estimator.
If the `BatchSize` variable is set to 0, it defaults to `EVM.RPCDefaultBatchSize`.

### BlockHistorySize
```toml
BlockHistorySize = 8 # Default
```
BlockHistorySize controls the number of past blocks to keep in memory to use as a basis for calculating a percentile gas price.

### CheckInclusionBlocks
```toml
CheckInclusionBlocks = 12 # Default
```
CheckInclusionBlocks is the number of recent blocks to use to detect if there is a transaction propagation/connectivity issue, and to prevent bumping in these cases.
This can help avoid the situation where RPC nodes are not propagating transactions for some non-price-related reason (e.g. go-ethereum bug, networking issue etc) and bumping gas would not help.

Set to zero to disable connectivity checking completely.

### CheckInclusionPercentile
```toml
CheckInclusionPercentile = 90 # Default
```
CheckInclusionPercentile controls the percentile that a transaction must have been higher than for all the blocks in the inclusion check window in order to register as a connectivity issue.

For example, if CheckInclusionBlocks=12 and CheckInclusionPercentile=90 then further bumping will be prevented for any transaction with any attempt that has a higher price than the 90th percentile for the most recent 12 blocks.

### EIP1559FeeCapBufferBlocks
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
EIP1559FeeCapBufferBlocks = 13 # Example
```
EIP1559FeeCapBufferBlocks controls the buffer blocks to add to the current base fee when sending a transaction. By default, the gas bumping threshold + 1 block is used.

(Only applies to EIP-1559 transactions)

### TransactionPercentile
```toml
TransactionPercentile = 60 # Default
```
TransactionPercentile specifies gas price to choose. E.g. if the block history contains four transactions with gas prices `[100, 200, 300, 400]` then picking 25 for this number will give a value of 200. If the calculated gas price is higher than `GasPriceDefault` then the higher price will be used as the base price for new transactions.

Must be in range 0-100.

Only has an effect if gas updater is enabled.

Think of this number as an indicator of how aggressive you want your node to price its transactions.

Setting this number higher will cause the Chainlink node to select higher gas prices.

Setting it lower will tend to set lower gas prices.

## GasEstimator.FeeHistory
```toml
[GasEstimator.FeeHistory]
CacheTimeout = '10s' # Default
```


### CacheTimeout
```toml
CacheTimeout = '10s' # Default
```
CacheTimeout is the time to wait in order to refresh the cached values stored in the FeeHistory estimator. A small jitter is applied so the timeout won't be exactly the same each time.

You want this value to be close to the block time. For slower chains, like Ethereum, you can set it to 12s, the same as the block time. For faster chains you can skip a block or two
and set it to two times the block time i.e. on Optimism you can set it to 4s. Ideally, you don't want to go lower than 1s since the RTT times of the RPC requests will be comparable to
the timeout. The estimator is already adding a buffer to account for a potential increase in prices within one or two blocks. On the other hand, slower frequency will fail to refresh
the prices and end up in stale values.

## HeadTracker
```toml
[HeadTracker]
HistoryDepth = 100 # Default
MaxBufferSize = 3 # Default
SamplingInterval = '1s' # Default
FinalityTagBypass = true # Default
MaxAllowedFinalityDepth = 10000 # Default
PersistenceEnabled = true # Default
```
The head tracker continually listens for new heads from the chain.

In addition to these settings, it log warnings if `EVM.NoNewHeadsThreshold` is exceeded without any new blocks being emitted.

### HistoryDepth
```toml
HistoryDepth = 100 # Default
```
HistoryDepth tracks the top N blocks on top of the latest finalized block to keep in the `heads` database table.
Note that this can easily result in MORE than `N + finality depth`  records since in the case of re-orgs we keep multiple heads for a particular block height.
Higher values help reduce number of RPC requests performed by TXM's Finalizer and improve TXM's Confirmer reorg protection on restarts.
At the same time, setting the value too high could lead to higher CPU consumption. The following formula could be used to calculate the optimal value: `expected_downtime_on_restart/block_time`.

### MaxBufferSize
```toml
MaxBufferSize = 3 # Default
```
MaxBufferSize is the maximum number of heads that may be
buffered in front of the head tracker before older heads start to be
dropped. You may think of it as something like the maximum permittable "lag"
for the head tracker before we start dropping heads to keep up.

### SamplingInterval
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
SamplingInterval = '1s' # Default
```
SamplingInterval means that head tracker callbacks will at maximum be made once in every window of this duration. This is a performance optimisation for fast chains. Set to 0 to disable sampling entirely.

### FinalityTagBypass
```toml
FinalityTagBypass = true # Default
```
FinalityTagBypass disables FinalityTag support in HeadTracker and makes it track blocks up to FinalityDepth from the most recent head.
It should only be used on chains with an extremely large actual finality depth (the number of blocks between the most recent head and the latest finalized block).
Has no effect if `FinalityTagsEnabled` = false

### MaxAllowedFinalityDepth
```toml
MaxAllowedFinalityDepth = 10000 # Default
```
MaxAllowedFinalityDepth - defines maximum number of blocks between the most recent head and the latest finalized block.
If actual finality depth exceeds this number, HeadTracker aborts backfill and returns an error.
Has no effect if `FinalityTagsEnabled` = false

### PersistenceEnabled
```toml
PersistenceEnabled = true # Default
```
PersistenceEnabled defines whether HeadTracker needs to store heads in the database.
Persistence is helpful on chains with large finality depth, where fetching blocks from the latest to the latest finalized takes a lot of time.
On chains with fast finality, the persistence layer does not improve the chain's load time and only consumes database resources (mainly IO).
NOTE: persistence should not be disabled for products that use LogBroadcaster, as it might lead to missed on-chain events.

## KeySpecific
```toml
[[KeySpecific]]
Key = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
GasEstimator.PriceMax = '79 gwei' # Example
```


### Key
```toml
Key = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
```
Key is the account to apply these settings to

### PriceMax
```toml
GasEstimator.PriceMax = '79 gwei' # Example
```
GasEstimator.PriceMax overrides the maximum gas price for this key. See EVM.GasEstimator.PriceMax.

## NodePool
```toml
[NodePool]
PollFailureThreshold = 5 # Default
PollInterval = '10s' # Default
SelectionMode = 'HighestHead' # Default
SyncThreshold = 5 # Default
LeaseDuration = '0s' # Default
NodeIsSyncingEnabled = false # Default
FinalizedBlockPollInterval = '5s' # Default
EnforceRepeatableRead = true # Default
DeathDeclarationDelay = '1m' # Default
NewHeadsPollInterval = '0s' # Default
VerifyChainID = true # Default
```
The node pool manages multiple RPC endpoints.

In addition to these settings, `EVM.NoNewHeadsThreshold` controls how long to wait after receiving no new heads before marking the node as out-of-sync.

### PollFailureThreshold
```toml
PollFailureThreshold = 5 # Default
```
PollFailureThreshold indicates how many consecutive polls must fail in order to mark a node as unreachable.

Set to zero to disable poll checking.

### PollInterval
```toml
PollInterval = '10s' # Default
```
PollInterval controls how often to poll the node to check for liveness.

Set to zero to disable poll checking.

### SelectionMode
```toml
SelectionMode = 'HighestHead' # Default
```
SelectionMode controls node selection strategy:
- HighestHead: use the node with the highest head number
- RoundRobin: rotate through nodes, per-request
- PriorityLevel: use the node with the smallest order number
- TotalDifficulty: use the node with the greatest total difficulty

### SyncThreshold
```toml
SyncThreshold = 5 # Default
```
SyncThreshold controls how far a node may lag behind the best node before being marked out-of-sync.
Depending on `SelectionMode`, this represents a difference in the number of blocks (`HighestHead`, `RoundRobin`, `PriorityLevel`), or total difficulty (`TotalDifficulty`).

Set to 0 to disable this check.

### LeaseDuration
```toml
LeaseDuration = '0s' # Default
```
LeaseDuration is the minimum duration that the selected "best" node (as defined by SelectionMode) will be used,
before switching to a better one if available. It also controls how often the lease check is done.
Setting this to a low value (under 1m) might cause RPC to switch too aggressively.
Recommended value is over 5m

Set to '0s' to disable

### NodeIsSyncingEnabled
```toml
NodeIsSyncingEnabled = false # Default
```
NodeIsSyncingEnabled is a flag that enables `syncing` health check on each reconnection to an RPC.
Node transitions and remains in `Syncing` state while RPC signals this state (In case of Ethereum `eth_syncing` returns anything other than false).
All of the requests to node in state `Syncing` are rejected.

Set true to enable this check

### FinalizedBlockPollInterval
```toml
FinalizedBlockPollInterval = '5s' # Default
```
FinalizedBlockPollInterval controls how often to poll RPC for new finalized blocks.
The finalized block is only used to report to the `pool_rpc_node_highest_finalized_block` metric. We plan to use it
in RPCs health assessment in the future.
If `FinalityTagEnabled = false`, poll is not performed and `pool_rpc_node_highest_finalized_block` is
reported based on latest block and finality depth.

Set to 0 to disable.

### EnforceRepeatableRead
```toml
EnforceRepeatableRead = true # Default
```
EnforceRepeatableRead defines if Core should only use RPCs whose most recently finalized block is greater or equal to
`highest finalized block - FinalizedBlockOffset`. In other words, exclude RPCs lagging on latest finalized
block.

Set false to disable

### DeathDeclarationDelay
```toml
DeathDeclarationDelay = '1m' # Default
```
DeathDeclarationDelay defines the minimum duration an RPC must be in unhealthy state before producing an error log message.
Larger values might be helpful to reduce the noisiness of health checks like `EnforceRepeatableRead = true', which might be falsely
trigger declaration of `FinalizedBlockOutOfSync` due to insignificant network delays in broadcasting of the finalized state among RPCs.
Should be greater than `FinalizedBlockPollInterval`.
Unhealthy RPC will not be picked to handle a request even if this option is set to a nonzero value.

### NewHeadsPollInterval
```toml
NewHeadsPollInterval = '0s' # Default
```
NewHeadsPollInterval define an interval for polling new block periodically using http client rather than subscribe to ws feed

Set to 0 to disable.

### VerifyChainID
```toml
VerifyChainID = true # Default
```
VerifyChainID enforces RPC Client ChainIDs to match configured ChainID

## NodePool.Errors
:warning: **_ADVANCED_**: _Do not change these settings unless you know what you are doing._
```toml
[NodePool.Errors]
NonceTooLow = '(: |^)nonce too low' # Example
NonceTooHigh = '(: |^)nonce too high' # Example
ReplacementTransactionUnderpriced = '(: |^)replacement transaction underpriced' # Example
LimitReached = '(: |^)limit reached' # Example
TransactionAlreadyInMempool = '(: |^)transaction already in mempool' # Example
TerminallyUnderpriced = '(: |^)terminally underpriced' # Example
InsufficientEth = '(: |^)insufficeint eth' # Example
TxFeeExceedsCap = '(: |^)tx fee exceeds cap' # Example
L2FeeTooLow = '(: |^)l2 fee too low' # Example
L2FeeTooHigh = '(: |^)l2 fee too high' # Example
L2Full = '(: |^)l2 full' # Example
TransactionAlreadyMined = '(: |^)transaction already mined' # Example
Fatal = '(: |^)fatal' # Example
ServiceUnavailable = '(: |^)service unavailable' # Example
TooManyResults = '(: |^)too many results' # Example
```
Errors enable the node to provide custom regex patterns to match against error messages from RPCs.

### NonceTooLow
```toml
NonceTooLow = '(: |^)nonce too low' # Example
```
NonceTooLow is a regex pattern to match against nonce too low errors.

### NonceTooHigh
```toml
NonceTooHigh = '(: |^)nonce too high' # Example
```
NonceTooHigh is a regex pattern to match against nonce too high errors.

### ReplacementTransactionUnderpriced
```toml
ReplacementTransactionUnderpriced = '(: |^)replacement transaction underpriced' # Example
```
ReplacementTransactionUnderpriced is a regex pattern to match against replacement transaction underpriced errors.

### LimitReached
```toml
LimitReached = '(: |^)limit reached' # Example
```
LimitReached is a regex pattern to match against limit reached errors.

### TransactionAlreadyInMempool
```toml
TransactionAlreadyInMempool = '(: |^)transaction already in mempool' # Example
```
TransactionAlreadyInMempool is a regex pattern to match against transaction already in mempool errors.

### TerminallyUnderpriced
```toml
TerminallyUnderpriced = '(: |^)terminally underpriced' # Example
```
TerminallyUnderpriced is a regex pattern to match against terminally underpriced errors.

### InsufficientEth
```toml
InsufficientEth = '(: |^)insufficeint eth' # Example
```
InsufficientEth is a regex pattern to match against insufficient eth errors.

### TxFeeExceedsCap
```toml
TxFeeExceedsCap = '(: |^)tx fee exceeds cap' # Example
```
TxFeeExceedsCap is a regex pattern to match against tx fee exceeds cap errors.

### L2FeeTooLow
```toml
L2FeeTooLow = '(: |^)l2 fee too low' # Example
```
L2FeeTooLow is a regex pattern to match against l2 fee too low errors.

### L2FeeTooHigh
```toml
L2FeeTooHigh = '(: |^)l2 fee too high' # Example
```
L2FeeTooHigh is a regex pattern to match against l2 fee too high errors.

### L2Full
```toml
L2Full = '(: |^)l2 full' # Example
```
L2Full is a regex pattern to match against l2 full errors.

### TransactionAlreadyMined
```toml
TransactionAlreadyMined = '(: |^)transaction already mined' # Example
```
TransactionAlreadyMined is a regex pattern to match against transaction already mined errors.

### Fatal
```toml
Fatal = '(: |^)fatal' # Example
```
Fatal is a regex pattern to match against fatal errors.

### ServiceUnavailable
```toml
ServiceUnavailable = '(: |^)service unavailable' # Example
```
ServiceUnavailable is a regex pattern to match against service unavailable errors.

### TooManyResults
```toml
TooManyResults = '(: |^)too many results' # Example
```
TooManyResults is a regex pattern to match an eth_getLogs error indicating the result set is too large to return

## OCR
```toml
[OCR]
ContractConfirmations = 4 # Default
ContractTransmitterTransmitTimeout = '10s' # Default
DatabaseTimeout = '10s' # Default
DeltaCOverride = "168h" # Default
DeltaCJitterOverride = "1h" # Default
ObservationGracePeriod = '1s' # Default
```


### ContractConfirmations
```toml
ContractConfirmations = 4 # Default
```
ContractConfirmations sets `OCR.ContractConfirmations` for this EVM chain.

### ContractTransmitterTransmitTimeout
```toml
ContractTransmitterTransmitTimeout = '10s' # Default
```
ContractTransmitterTransmitTimeout sets `OCR.ContractTransmitterTransmitTimeout` for this EVM chain.

### DatabaseTimeout
```toml
DatabaseTimeout = '10s' # Default
```
DatabaseTimeout sets `OCR.DatabaseTimeout` for this EVM chain.

### DeltaCOverride
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
DeltaCOverride = "168h" # Default
```
DeltaCOverride (and `DeltaCJitterOverride`) determine the config override DeltaC.
DeltaC is the maximum age of the latest report in the contract. If the maximum age is exceeded, a new report will be
created by the report generation protocol.

### DeltaCJitterOverride
:warning: **_ADVANCED_**: _Do not change this setting unless you know what you are doing._
```toml
DeltaCJitterOverride = "1h" # Default
```
DeltaCJitterOverride is the range for jitter to add to `DeltaCOverride`.

### ObservationGracePeriod
```toml
ObservationGracePeriod = '1s' # Default
```
ObservationGracePeriod sets `OCR.ObservationGracePeriod` for this EVM chain.

## Nodes
```toml
[[Nodes]]
Name = 'foo' # Example
WSURL = 'wss://web.socket/test' # Example
HTTPURL = 'https://foo.web' # Example
HTTPURLExtraWrite = 'https://foo.web/extra' # Example
SendOnly = false # Default
Order = 100 # Default
```


### Name
```toml
Name = 'foo' # Example
```
Name is a unique (per-chain) identifier for this node.

### WSURL
```toml
WSURL = 'wss://web.socket/test' # Example
```
WSURL is the WS(S) endpoint for this node. Required for primary nodes when `LogBroadcasterEnabled` is `true`

### HTTPURL
```toml
HTTPURL = 'https://foo.web' # Example
```
HTTPURL is the HTTP(S) endpoint for this node. Required for all nodes.

### HTTPURLExtraWrite
```toml
HTTPURLExtraWrite = 'https://foo.web/extra' # Example
```
HTTPURLExtraWrite is the HTTP(S) endpoint used for chains that require a separate endpoint for writing on-chain.

### SendOnly
```toml
SendOnly = false # Default
```
SendOnly limits usage to sending transaction broadcasts only. With this enabled, only HTTPURL is required, and WSURL is not used.

### Order
```toml
Order = 100 # Default
```
Order of the node in the pool, will takes effect if `SelectionMode` is `PriorityLevel` or will be used as a tie-breaker for `HighestHead` and `TotalDifficulty`

## OCR2.Automation
```toml
[OCR2.Automation]
GasLimit = 5400000 # Default
```


### GasLimit
```toml
GasLimit = 5400000 # Default
```
GasLimit controls the gas limit for transmit transactions from ocr2automation job.

## Workflow
```toml
[Workflow]
FromAddress = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
ForwarderAddress = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
GasLimitDefault = 400_000 # Default
```


### FromAddress
```toml
FromAddress = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
```
FromAddress is Address of the transmitter key to use for workflow writes.

### ForwarderAddress
```toml
ForwarderAddress = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292' # Example
```
ForwarderAddress is the keystone forwarder contract address on chain.

### GasLimitDefault
```toml
GasLimitDefault = 400_000 # Default
```
GasLimitDefault is the default gas limit for workflow transactions.

