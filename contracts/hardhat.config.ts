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
  },
  solidity: {
    compilers: [
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
  },
}

export default config
