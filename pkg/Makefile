.PHONY: testdb
testdb:
	cd $(shell go mod download -json github.com/smartcontractkit/chainlink/v2@4cb89093788a1c3acfc183fb86adbd722ebb6bfb | jq -r .Dir) && \
	go run ./core/store/cmd/preparetest
