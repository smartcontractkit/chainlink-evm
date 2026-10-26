.PHONY: testdb
testdb:
	cd $(shell go mod download -json github.com/smartcontractkit/chainlink/v2@491bf1b63c64bf509cd86f18cddef89d7c9fac2a | jq -r .Dir) && \
	go run ./core/store/cmd/preparetest
	psql "$(shell printf '%s' "$(CL_DATABASE_URL)" | sed -E 's@/[^/?]+(\?.*)?$$@/postgres\1@')" \
		-c 'DROP DATABASE IF EXISTS chainlink_test_pristine WITH (FORCE)' \
		-c 'CREATE DATABASE chainlink_test_pristine WITH TEMPLATE chainlink_test'
