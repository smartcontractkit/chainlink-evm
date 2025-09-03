import { HardhatUserConfig } from 'hardhat/config'

const COMPILER_SETTINGS = {
  optimizer: {
    enabled: true,
    runs: 1000000,
  },
  metadata: {
    bytecodeHash: 'none',
  },
}

const config: HardhatUserConfig = {
  paths: {
    artifacts: './artifacts',
    cache: './cache',
    sources: './src/v0.8/automation',
    tests: './test',
  },
  solidity: {
    compilers: [
      {
        version: '0.8.16',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.19',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.26',
        settings: {
          ...COMPILER_SETTINGS,
          evmVersion: 'paris',
        },
      },
    ],
    overrides: {
      'src/v0.8/automation/AutomationForwarderLogic.sol': {
        version: '0.8.19',
        settings: COMPILER_SETTINGS,
      },
    },
  },
}

export default config
