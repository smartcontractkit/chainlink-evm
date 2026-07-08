.PHONY: testdb
testdb:
	cd $(shell go mod download -json github.com/smartcontractkit/chainlink/v2@ec649c7da3267b6013009dc612472ae702f61873 | jq -r .Dir) && \
	go run ./core/store/cmd/preparetest
