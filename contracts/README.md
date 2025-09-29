# Chainlink Smart Contracts

> [!IMPORTANT]
> Since `v1.5.0` of the Chainlink contracts package, some dependencies are no longer vendored and require the setup of remappings.
> See the setup instructions for use in Solidity projects.
>
> Since `v1.4.0` of the Chainlink contracts, the contracts have been moved to their own repository:
> [chainlink-evm](https://github.com/smartcontractkit/chainlink-evm).
> Prior to that, the contracts were part of the [main Chainlink repository](https://github.com/smartcontractkit/chainlink)

## Table of Contents

- [Installation](#installation)
- [Setup](#setup)
    - [Foundry](#foundry)
    - [Hardhat 3](#hardhat-3)
    - [Hardhat 2](#hardhat-2)
    - [Remix](#remix)
- [Package Directory Structure](#package-directory-structure)
- [Usage](#usage)
- [Local Development](#local-development)
- [Contributing](#contributing)
    - [Changesets](#changesets)


## Installation

> [!WARNING]  
> For use in Solidity project(s), see the setup instructions below.

```sh
# pnpm
$ pnpm add @chainlink/contracts
```

```sh
# npm
$ npm install @chainlink/contracts --save
```

## Setup

For use in Solidity projects, further configuration is required. 

This package relies on [remappings](https://docs.soliditylang.org/en/latest/path-resolution.html#import-remapping) to resolve import paths within your Solidity project(s). Each tool may handle [remappings](https://docs.soliditylang.org/en/latest/path-resolution.html#import-remapping) in a different manner. In the sections below, you will find detailed instructions on this process for popular tools.

Refer to the [Solidity remapping documentation](https://docs.soliditylang.org/en/latest/path-resolution.html#import-remapping) for more information.

<details id="foundry">
<summary>Foundry</summary>

### Step 1: Install the package

For use in your Foundry project, it is recommended to utilize `npm` or `pnpm` as your package manager for the use of this package instead of `forge install`. 

```sh
# pnpm
$ pnpm add @chainlink/contracts
```

```sh
# npm
$ npm install @chainlink/contracts --save
```

If you wish to utilize `forge install`, please see the [Foundry starter kit](https://github.com/smartcontractkit/foundry-starter-kit).

### Step 2: Define the external library directory

As we are using `npm`/`pnpm` as our package manager, define `node_modules` as an external library directory. This ensures Foundry recognizes dependencies installed via `npm`/`pnpm`.

In your project's `foundry.toml`, update the libs array to include the `node_modules` directory.

```
libs = ["lib", "node_modules"]
```

### Step 3: Setup remappings

Setup your project's remappings. See the [Foundry documentation](https://getfoundry.sh/guides/project-setup/dependencies#remapping-dependencies) for more information.

[Foundry](https://getfoundry.sh/guides/project-setup/project-layout#project-layout) consumes a `remappings.txt` file from the project root. Create or update `remappings.txt` with the following:

```
@chainlink/=node_modules/@chainlink
@openzeppelin/contracts@4.7.3=node_modules/@openzeppelin/contracts-4.7.3
@openzeppelin/contracts@4.8.3=node_modules/@openzeppelin/contracts-4.8.3
@openzeppelin/contracts@4.9.6=node_modules/@openzeppelin/contracts-4.9.6
@openzeppelin/contracts@5.0.2=node_modules/@openzeppelin/contracts-5.0.2
@openzeppelin/contracts@5.1.0=node_modules/@openzeppelin/contracts-5.1.0
@openzeppelin/contracts-upgradeable/=node_modules/@openzeppelin/contracts-upgradeable/
@arbitrum/=node_modules/@arbitrum/
@eth-optimism/=node_modules/@eth-optimism/
@scroll-tech/=node_modules/@scroll-tech/
@zksync/=node_modules/@zksync/
```

If your compilation reports unresolved imports from dependencies, add the corresponding additional remappings to `remappings.txt` (the format is `<prefix>=<resolved-path>/`).

### Step 4: Usage

```solidity
import {IVerifier} from '@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol';
```

See the [Foundry starter kit](https://github.com/smartcontractkit/foundry-starter-kit) for working examples.

</details>

<details id="hardhat-3">
<summary>Hardhat 3</summary>

### Step 1: Install the package

```sh
# pnpm
$ pnpm add @chainlink/contracts
```

```sh
# npm
$ npm install @chainlink/contracts --save
```

Hardhat 3 supports `remappings.txt` files in your project, as well as in git submodules and npm dependencies. Each `remappings.txt` file applies to the directory where it's located and all its subdirectories, similar to how .gitignore works. Hardhat 3 will utilize the `remappings.txt` file located within the root directory of the Chainlink contracts package.

### Step 2: Usage

```solidity
import {IVerifier} from '@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol';
```

See the [Hardhat 3 starter kit](https://github.com/smartcontractkit/hardhat-starter-kit/) for working examples.

</details>


<details id="hardhat-2">
<summary>Hardhat 2</summary>

### Step 1: Install the package

```sh
# pnpm
$ pnpm add @chainlink/contracts
```

```sh
# npm
$ npm install @chainlink/contracts --save
```

### Step 2: Setup remappings

Hardhat 2 does not read `remappings.txt` natively as seen in Foundry/Hardhat 3. To remap import paths, you may use a preprocessor that handles this at compile time. To see remapping examples in Hardhat 2, review the [Hardhat 2 starter kit](https://github.com/smartcontractkit/hardhat-starter-kit/).

### Step 3: Usage

```solidity
import {IVerifier} from '@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol';
```

</details>


<details id="remix">
<summary>Remix</summary>

Remix works out of the box and requires no additional setup or installation. The imported dependencies will be automatically installed.

```solidity
import {IVerifier} from '@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol';
```

</details>


## Package Directory Structure

> [!IMPORTANT]
> Since v1.5.0 of the Chainlink contracts, ABI files have been reorganized into subdirectories.
> Additionally, ABI files now follow a slightly updated naming scheme.

```sh
@chainlink/contracts
├── src # Solidity contracts
│   └── v0.8
└── abi # ABI JSON output
    └── v0.8
```

## Usage

The Solidity smart contracts themselves can be imported via the `src` directory of `@chainlink/contracts`:

```
@chainlink/contracts/src/v0.8/llo-feeds/v0.5.0/interfaces/IVerifier.sol
```

The ABI files themselves can be imported via the `abi` directory of `@chainlink/contracts`:

```
@chainlink/contracts/abi/v0.8/VRF/VRFCoordinatorV2_5.abi.json
```


## Local Development

**Note:** Contracts in `dev/` directories or with a typeAndVersion ending in `-dev` are under active development and are likely unaudited. Please refrain from using these in production applications.

```bash
# Clone Chainlink repository
$ git clone https://github.com/smartcontractkit/chainlink.git
$ cd contracts/
$ pnpm
```

Each Chainlink project has its own directory under `src/` which can be targeted using Foundry profiles. To test a specific project, run:

```bash
# Replace <project> with the product you want to test
export FOUNDRY_PROFILE=<project>
forge test
```

To test the llo-feeds (data streams) project:

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

Every PR that modifies any configuration or code should most likely be accompanied by a changeset file.

To install `changesets`:

1. Install `pnpm` if it is not already installed - [docs](https://pnpm.io/installation).
2. Run `pnpm install`.

Either after or before you create a commit, run the `pnpm changeset` command in the `contracts` directory to create an accompanying changeset entry which will reflect on the CHANGELOG for the next release.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
