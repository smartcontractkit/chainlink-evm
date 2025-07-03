## License

This repository contains two separate license regimes:

1. **LGPL-3.0-or-later** for all code in `./abigen` (the forked go-ethereum abigen).  
   See the full text in `LICENSE` under “GNU LESSER…”  
2. **MIT** for everything else in this repo.  
   See the full text in `LICENSE` under “MIT License”.


# CRE Generated Bindings (MVP)

This project utilizes a forked version of `abigen` (from go-ethereum)
that lets you generate Go bindings for your smart contracts using a custom template.

## Prerequisites

1. **Go**
   Install Go 1.18 or later:
   ```bash
   brew install go          # macOS (Homebrew)
   sudo apt install golang  # Ubuntu/Debian
   ```
2. **Solidity compiler**
   Install `solc` to compile or verify your contracts:
   ```bash
   npm install -g solc      # via npm
   brew install solidity    # macOS (Homebrew)
   ```

## Usage
### Programmatic API

```go
import "github.com/smartcontractkit/chainlink-evm/pkg/bindings"

func main() {
  err := bindings.GenerateBindings(
    "./pkg/bindings/build/MyContract_combined.json", // or "" if using abiPath
    "./pkg/bindings/MyContract.abi",                 // or "" for combined-json mode
    "bindings",                                       // Go package name
    "MyContract",                                     // typeName (single-ABI only)
    "./pkg/bindings/build/bindings.go",               // output file
  )
  if err != nil {
    log.Fatalf("generate bindings: %v", err)
  }
}
```