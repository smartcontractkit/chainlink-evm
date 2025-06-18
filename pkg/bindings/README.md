# CRE Generated Bindings (MVP)

## Prerequisites:
1. Install `go`
2. Install `solidity`

## Setup instructions:

1. Clone `geth` fork and checkout development branch
```
git clone https://github.com/pablolagreca/go-ethereum.git && git checkout abigen-configurable-template
```

2. Build custom abigen binary.
```
go build -o "$(go env GOPATH)/bin/abigen-cre" ./cmd/abigen
```

3. Generate the abi combined json file 
```
solc --combined-json abi,bin,metadata {SOLIDITY_FILE_LOCATION} > {ABI_FILE_LOCATION}
```

4. Run abigen-cre, passing in the cre source template:
```
abigen-cre --v2 \
  --abi {ABI_FILE_LOCATION} \
  --pkg bindings \
  --template ./sourcecre.go.tpl \
  --out ./build/bindings.go
```



