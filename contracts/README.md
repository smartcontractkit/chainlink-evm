# Chainlink Smart Contracts

> [!IMPORTANT] Since v1.4.0 of the Chainlink contracts, the contracts have been moved to their own repository:
> [chainlink-evm](https://github.com/smartcontractkit/chainlink-evm). 
> Prior to that, the contracts were part of the [main Chainlink repository](https://github.com/smartcontractkit/chainlink)

## Installation

#### Foundry (git)

> [!WARNING]
> When installing via git, the ref defaults to master when no tag is given.


```sh
$ forge install smartcontractkit/chainlink-evm@<version_tag>
```

Add `@chainlink/contracts/=lib/smartcontractkit/chainlink-evm/contracts/` in remappings.txt.

#### NPM
```sh
# pnpm
$ pnpm add @chainlink/contracts
```

```sh
# npm
$ npm install @chainlink/contracts --save
```

Add `@chainlink/contracts/=node_modules/@chainlink/contracts/` in remappings.txt.



### Directory Structure

```sh
@chainlink/contracts
├── src # Solidity contracts
│   └── v0.8
└── abi # ABI json output
    └── v0.8
```

### Usage

The solidity smart contracts themselves can be imported via the `src` directory of `@chainlink/contracts`:

```solidity
import {IVerifier} from '@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol';
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
$ git clone https://github.com/smartcontractkit/chainlink.git
$ cd contracts/
$ pnpm
```

Each Chainlink project has its own directory under `src/` which can be targeted using Foundry profiles.
To test a specific project, run:

```bash
# Replace <project> with the product you want to test
export FOUNDRY_PROFILE=<project>
forge test
```

To test the llo-feeds (data steams) project:

```bash
export FOUNDRY_PROFILE=llo-feeds
forge test
```

## Contributing

Please adhere to the [Solidity Style Guide](https://github.com/smartcontractkit/chainlink-evm/blob/develop/contracts/STYLE_GUIDE.md).

Contributions are welcome! Please refer to
[Chainlink's contributing guidelines](https://github.com/smartcontractkit/chainlink/blob/develop/docs/CONTRIBUTING.md) for detailed
contribution information.

Thank you!

### Changesets

We use [changesets](https://github.com/changesets/changesets) to manage versioning the contracts.

Every PR that modifies any configuration or code, should most likely accompanied by a changeset file.

To install `changesets`:
  1. Install `pnpm` if it is not already installed - [docs](https://pnpm.io/installation).
  2. Run `pnpm install`.

Either after or before you create a commit, run the `pnpm changeset` command in the `contracts` directory to create an accompanying changeset entry which will reflect on the CHANGELOG for the next release.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

