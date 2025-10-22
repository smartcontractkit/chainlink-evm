The offchainaggregator bindings have been added in a special way and this folder just preserves it:
 * On Feb 24, 2021 the OCR one was added https://github.com/smartcontractkit/chainlink/commit/27dc2451cff8ea033eafa94cad6b6f4406bc67f4
and then its ABI was frozen https://github.com/smartcontractkit/chainlink/commit/6cfeb1473b48719fc1a20ed83b9b67d236ba0ca9
 * On Dec 2, 2021 a newer version of the binding (OCR2) one was added https://github.com/smartcontractkit/chainlink/commit/783112d5816a1ab94569e3eef5e641716871d121
already in a frozen state (presumably compiled from a snapshot of libocr's `contract2/OCR2Aggregator.sol`)

Note that two contracts are **different** (e.g., their topic hashes are different). In particular, the later one is an updated version of the former one - https://github.com/smartcontractkit/libocr/commit/16d15bf6fb4408102828cc1faaf6ff8f1afdee25.