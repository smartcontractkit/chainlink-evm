.PHONY: testdb
testdb:
	cd $(shell go mod download -json github.com/smartcontractkit/chainlink/v2@a4276acb9b6cc235b9d101f093c619bb3ca5f4d6 | jq -r .Dir) && \
	go run ./core/store/cmd/preparetest
	psql "$(shell printf '%s' "$(CL_DATABASE_URL)" | sed -E 's@/[^/?]+(\?.*)?$$@/postgres\1@')" \
		-c 'DROP DATABASE IF EXISTS chainlink_test_pristine WITH (FORCE)' \
		-c 'CREATE DATABASE chainlink_test_pristine WITH TEMPLATE chainlink_test'
