# TXM Integration Tests

This directory contains integration tests for the Transaction Manager v2 (TXMv2) package. These tests verify different end-to-end scenarios, including transaction creation, queuing, processing, error handling, etc. More tests can be added by using the builder methods.

## Test Modes

The tests support different simulation modes to exercise various TXM behaviors. The modes falsely return successful messages when broadcasting transactions in order to inject nonce gaps and transactions that
are never going to be mined unless the TXM takes an action.

- **Standard**: Normal transaction flow without any special error conditions
- **Retransmission**: Simulates scenarios where transactions appear successful but weren't actually mined
- **StuckTxDetection**: Tests the stuck transaction detection and recovery mechanism
- **ErrorHandling**: Validates error handling logic, particularly for nonce reassignment scenarios


## Test Functions

### `TestIntegration_StandardFlow`

Tests the standard transaction flow by creating multiple transaction requests and verifying they are all processed successfully. Given a number of transactions larger than MaxInFlightSubset will
trigger a throttling scenario and ensure the TXM handles them gracefully.

### `TestIntegration_Retransmission`

Tests the TXM's retransmission logic by simulating a scenario where every other transaction attempt appears successful but wasn't actually mined, requiring the TXM to handle retransmission.

### `TestIntegration_StuckTxDetection`

Tests the TXM's stuck transaction detection by injecting a mix of stuck and non-stuck transactions and verifying the detection mechanism works correctly.

### `TestIntegration_ErrorHandling`

Tests the TXM's error handling logic by injecting transactions with specific MetaClient error messages and verifying that nonce reassignment is handled correctly.

## Running the Tests

### Default Mode (DEVNET - Simulated Backend)

By default, the tests run against a simulated Ethereum backend. No additional configuration is required:

```bash
go test -v ./pkg/txm/integration-tests/...
```

### Testnet Mode

To run tests against a real testnet, set the `ENV` environment variable to `TESTNET`. Testnet needs to be EVM-based:

```bash
ENV=TESTNET go test -v ./pkg/txm/integration-tests/...
```

Copy the `env-example.toml` and `configs-example.toml` files and rename them. You must provide the following values:

1. **`env.toml`**: Contains testnet-specific environment variables:
   - `RPC`: The RPC endpoint URL for the testnet
   - `PrivateKey`: The private key (without 0x prefix) for the account to use

2. **`configs.toml`**: Contains test configuration:
   - `BlockTime`: Block time in seconds
   - `EIP1559DynamicFees`: Whether to use EIP-1559 dynamic fee transactions
   - `BumpThreshold`: Number of blocks to wait before retransmission. The name is used to be backwards compatible with TXMv1.

#### Example Configuration Files

**`env.toml`**:
```toml
RPC = 'https://your-testnet-rpc-endpoint'
PrivateKey = 'your-private-key-without-0x-prefix'
```

**`configs.toml`**:
```toml
BlockTime = 12
EIP1559DynamicFees = true
BumpThreshold = 3
```

## Test Architecture

### Setup Functions

- **`setupBackend`**: Main setup function that initializes the test environment based on the `ENV` variable
- **`setupDevnetTXM`**: Sets up TXM with a simulated backend (default)
- **`setupTestnetTXM`**: Sets up TXM with a real testnet connection
- **`setupGasEstimator`**: Configures and starts the gas estimator
- **`setupSimulatedBackendClient`**: Creates a simulated Ethereum backend with initial transaction history

### Components

The tests set up a complete TXM instance with:
- **Gas Estimator**: Estimates gas prices and fees
- **Keystore**: Manages private keys and signing
- **Storage**: In-memory transaction storage
- **Attempt Builder**: Builds transaction attempts with appropriate gas settings
- **Stuck TX Detector**: (Optional) Detects and handles stuck transactions
- **Error Handler**: (Optional) Handles transaction errors

## Test Flow

1. **Setup**: Initialize TXM with appropriate configuration and simulation mode
2. **Transaction Creation**: Create multiple transaction requests (default: 20)
3. **Trigger**: Manually trigger TXM processing (instead of waiting for the next cycle)
4. **Verification**: Wait until all transaction queues are empty, confirming all transactions were processed

## Configuration

### Default Configuration (DEVNET)

When running in DEVNET mode, the following defaults are used:
- Chain ID: 1337
- Block Time: Test interval
- EIP-1559: Enabled
- Bump Threshold: 3 blocks
- Gas Limit (Default): 30,000
- Gas Limit (Transfer): 21,000

*Note: transaction mining is executed when client.Commit() method is called. To emulate a blockchain behaviour, the method is called in each interval of `waitUntilQueuesAreEmpty`.*

### Testnet Configuration

Testnet configuration is loaded from `configs.toml` and `env.toml` files. See the example files in this directory for reference.

## Important Notes

- **Private Keys**: Never commit real private keys or testnet credentials to version control
- **Testnet Costs**: Running tests against a testnet will consume testnet ETH for gas fees. You need to fund your private key before running the testnet tests otherwise they will fail.
- **Simulated Backend**: The DEVNET mode uses a simulated backend, which is faster and doesn't require network access or funds

## Troubleshooting

### Tests Fail in TESTNET Mode

1. Verify your `env.toml` contains valid RPC endpoint and private key
2. Ensure the account has sufficient testnet ETH for gas fees
3. Check that the RPC endpoint is accessible and responsive
4. Verify the `configs.toml` settings are appropriate for your testnet

### Tests Timeout

- Increase the timeout in `waitUntilQueuesAreEmpty` if your testnet has slow block times. If you run all the test scenarios on a slow chain like Ethereum with a large number of transactions, the transaction confirmation will be rather slow.
- Verify that transactions are actually being mined on the testnet
