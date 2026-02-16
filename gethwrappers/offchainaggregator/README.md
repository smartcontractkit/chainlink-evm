The offchainaggregator bindings have been added in a special way and this folder just preserves it:

* **OCR** — On Feb 24, 2021 the OCR one was added https://github.com/smartcontractkit/chainlink/commit/27dc2451cff8ea033eafa94cad6b6f4406bc67f4 and then its ABI was frozen https://github.com/smartcontractkit/chainlink/commit/6cfeb1473b48719fc1a20ed83b9b67d236ba0ca9

* **OCR2** — On Dec 2, 2021 a newer version of the binding (OCR2) was added https://github.com/smartcontractkit/chainlink/commit/783112d5816a1ab94569e3eef5e641716871d121 already in a frozen state (presumably compiled from a snapshot of libocr's `contract2/OCR2Aggregator.sol`). Generated package: `generated/ocr2/offchainaggregator/`. Constructor takes 12 args (billing params + link, min/max, access controllers, decimals, description).

* **Data Feeds 1 (AccessControlledOCR2Aggregator)** — Binding for the Data Feeds 1 OCR2 aggregator with a 7-arg constructor (`_link`, `_minAnswer`, `_maxAnswer`, `_billingAccessController`, `_requesterAccessController`, `_decimals`, `description`). Generated package: `generated/ocr2/data-feeds-1-offchainaggregator/`. Source artifact: [gauntlet-evm AccessControlledOCR2Aggregator.json](https://github.com/smartcontractkit/gauntlet-evm/blob/develop/packages/evm-gauntlet-ocr/artifacts/evm/AccessControlledOCR2Aggregator.json). Local path in gauntlet-evm: `packages/evm-gauntlet-ocr/artifacts/evm/AccessControlledOCR2Aggregator.json`. To regenerate so the binding matches the artifact exactly: extract `.abi` to a JSON file and `.bytecode.object` to a hex file, then run `abigen -abi <abi.json> -bin <bin.txt> -pkg datafeeds1offchainaggregator -type OffchainAggregator -out gethwrappers/offchainaggregator/generated/ocr2/data-feeds-1-offchainaggregator/offchainaggregator.go`. This binding is used for deploy, setPayees, and other contract calls; **setConfig** calldata is not built by this generated binding and must be encoded externally (same layout as gauntlet-evm `encoding.ts`: version + int192 min + int192 max).

  **Verification that the binding matches the Gauntlet artifact:** From the chainlink-evm repo, run the ABI verification test with the artifact path (e.g. from a local gauntlet-evm clone):

  ```bash
  GAUNTLET_ARTIFACT_PATH=/path/to/gauntlet-evm/packages/evm-gauntlet-ocr/artifacts/evm/AccessControlledOCR2Aggregator.json go test -v ./gethwrappers/offchainaggregator/generated/ocr2/data-feeds-1-offchainaggregator/ -run TestGeneratedABIMatchesGauntletArtifact
  ```

  The test parses both ABIs and asserts that every method and event has the same selector/topic (ID). Without `GAUNTLET_ARTIFACT_PATH` the test is skipped. CI can run this when the artifact is available (e.g. checkout gauntlet-evm or download the artifact).

Note that the OCR and OCR2 contracts are **different** (e.g., their topic hashes are different). In particular, the latter one is an updated version of the former one — https://github.com/smartcontractkit/libocr/commit/16d15bf6fb4408102828cc1faaf6ff8f1afdee25.
