.PHONY: testdb
testdb:
	cd $(shell go mod download -json github.com/smartcontractkit/chainlink/v2@4a6fa896594e6817ebc36f64030770298b3cd939 | jq -r .Dir) && \
	go run ./core/store/cmd/preparetest
