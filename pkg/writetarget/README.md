# Generalized Write Target

This framework allows for any chain to implement the write target with minimal friction simply by implementing the interface [target_strategy](https://github.com/smartcontractkit/chainlink-framework/blob/0647c811e8e34635171517e64571650c59402d6a/capabilities/writetarget/write_target.go#L50-L57)

[Aptos Implementation](https://github.com/smartcontractkit/chainlink-aptos/blob/133766330253521cb0eb23b8d86ca89e187d5bc2/relayer/write_target/strategy.go#L1-L171)
[EVM Implementation](https://github.com/smartcontractkit/chainlink/blob/12c6f874df3ea9571eb3c8aa98f4a41c4d5b49b2/core/services/relay/evm/target_strategy.go#L1-L226)