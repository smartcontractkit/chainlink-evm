# Chainlink CRE Smart Contracts

## Installation

#### NPM
```sh
# pnpm
$ pnpm add @chainlink/contracts-cre
```

```sh
# npm
$ npm install @chainlink/contracts-cre --save
```

Add `@chainlink/contracts-cre/=node_modules/@chainlink/contracts-cre/` in remappings.txt.


### Usage

The solidity smart contracts themselves can be imported via the `src` directory of `@chainlink/contracts-cre`:

```solidity
import {CapabilitiesRegistry} from '@chainlink/contracts-cre/src/v2/CapabilitiesRegistry.sol';
```

### Remapping

This repository uses [Solidity remappings](https://docs.soliditylang.org/en/v0.8.20/using-the-compiler.html#compiler-remapping) to resolve imports.
The remapping is defined in the `remappings.txt` file.


## Local Development

Note:
Contracts in `dev/` directories or with a typeAndVersion ending in `-dev` are under active development
and are likely unaudited.
Please refrain from using these in production applications.

```bash
# Clone Chainlink repository
$ git clone https://github.com/smartcontractkit/chainlink-evm.git
$ cd contracts/cre/
$ pnpm i
```

To test the v2 contracts:

```bash
export FOUNDRY_PROFILE=v2
forge test
```