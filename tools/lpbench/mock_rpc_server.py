#!/usr/bin/env python3
"""Tiny deterministic Arbitrum-like JSON-RPC server for exercising the probe.

It models the two RPC calls that matter to the log poller:
  - eth_getBlockByNumber
  - eth_getLogs (by blockHash or by range)

Every method call sleeps --delay-ms.  A JSON-RPC batch is either processed
serially (--batch-concurrency 1) or with a thread pool, which is the variable
the real experiment is trying to measure.  --max-batch simulates a provider
batch-size limit.
"""

import argparse
import hashlib
import json
import threading
from concurrent.futures import ThreadPoolExecutor
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer


def h32(s: str) -> str:
    return "0x" + hashlib.sha256(s.encode()).hexdigest()


class Chain:
    def __init__(self, n_blocks: int = 200_000):
        self.n_blocks = n_blocks
        self.hashes = [h32(f"block-{i}") for i in range(n_blocks)]
        self.timestamps = [1_700_000_000 + i // 4 for i in range(n_blocks)]
        self.logs_by_hash = {}
        for i in range(3, n_blocks, 3):
            self.logs_by_hash[self.hashes[i]] = [{
                "address": "0x00000000000000000000000000000000000000aa",
                "topics": [h32("event"), h32(f"request-{i}")],
                "data": "0x" + f"{i:064x}",
                "blockNumber": hex(i),
                "blockHash": self.hashes[i],
                "transactionHash": h32(f"tx-{i}"),
                "transactionIndex": "0x0",
                "logIndex": "0x0",
                "removed": False,
            }]

    def block(self, n: int):
        if n < 0 or n >= self.n_blocks:
            return None
        return {
            "number": hex(n),
            "hash": self.hashes[n],
            "parentHash": self.hashes[n - 1] if n > 0 else "0x" + "00" * 32,
            "timestamp": hex(self.timestamps[n]),
            "transactions": [],
            # A realistic header has more fields; the probe only needs these.
        }


CHAIN = Chain()


class Handler(BaseHTTPRequestHandler):
    protocol_version = "HTTP/1.1"
    delay = 0.160
    batch_concurrency = 16
    max_batch = 1000

    def log_message(self, *args):
        pass

    def _one(self, req):
        method = req.get("method")
        params = req.get("params") or []
        if self.delay:
            # Simulate the provider-side cost of one method call.
            import time
            time.sleep(self.delay)

        if method == "eth_blockNumber":
            return {"jsonrpc": "2.0", "id": req.get("id"), "result": hex(CHAIN.n_blocks - 1)}

        if method == "eth_chainId":
            return {"jsonrpc": "2.0", "id": req.get("id"), "result": "0xa4b1"}

        if method == "eth_getBlockByNumber":
            try:
                n = int(params[0], 16) if isinstance(params[0], str) and params[0].startswith("0x") else int(params[0])
            except (ValueError, TypeError, IndexError):
                return {"jsonrpc": "2.0", "id": req.get("id"), "error": {"code": -32602, "message": "bad block number"}}
            b = CHAIN.block(n)
            return {"jsonrpc": "2.0", "id": req.get("id"), "result": b}

        if method == "eth_getLogs":
            f = params[0] if params else {}
            if "blockHash" in f:
                logs = CHAIN.logs_by_hash.get(f["blockHash"], [])
            else:
                from_n = int(f.get("fromBlock", "0x0"), 16)
                to_n = int(f.get("toBlock", hex(CHAIN.n_blocks - 1)), 16)
                logs = []
                for n in range(from_n, to_n + 1):
                    logs.extend(CHAIN.logs_by_hash.get(CHAIN.hashes[n], []))
            return {"jsonrpc": "2.0", "id": req.get("id"), "result": logs}

        return {"jsonrpc": "2.0", "id": req.get("id"), "error": {"code": -32601, "message": f"method not found: {method}"}}

    def do_POST(self):
        length = int(self.headers.get("Content-Length", "0"))
        raw = self.rfile.read(length)
        try:
            payload = json.loads(raw)
        except json.JSONDecodeError:
            self.send_error(400)
            return

        if isinstance(payload, list):
            if len(payload) > self.max_batch:
                body = json.dumps({
                    "jsonrpc": "2.0",
                    "id": None,
                    "error": {"code": -32600, "message": f"batch size {len(payload)} exceeds limit {self.max_batch}"},
                }).encode()
                self.send_response(413)
                self.send_header("Content-Type", "application/json")
                self.send_header("Content-Length", str(len(body)))
                self.end_headers()
                self.wfile.write(body)
                return

            if self.batch_concurrency > 1:
                with ThreadPoolExecutor(max_workers=self.batch_concurrency) as ex:
                    out = list(ex.map(self._one, payload))
            else:
                out = [self._one(r) for r in payload]
        else:
            out = self._one(payload)

        body = json.dumps(out).encode()
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(body)))
        self.end_headers()
        self.wfile.write(body)


def main():
    p = argparse.ArgumentParser()
    p.add_argument("--port", type=int, default=8545)
    p.add_argument("--delay-ms", type=float, default=160.0)
    p.add_argument("--batch-concurrency", type=int, default=16,
                   help="1 = serial provider-side batch processing; >1 = parallel")
    p.add_argument("--max-batch", type=int, default=1000)
    args = p.parse_args()

    Handler.delay = args.delay_ms / 1000.0
    Handler.batch_concurrency = args.batch_concurrency
    Handler.max_batch = args.max_batch
    srv = ThreadingHTTPServer(("127.0.0.1", args.port), Handler)
    print(f"mock RPC on http://127.0.0.1:{args.port} "
          f"delay={args.delay_ms}ms batch_concurrency={args.batch_concurrency} "
          f"max_batch={args.max_batch}", flush=True)
    srv.serve_forever()


if __name__ == "__main__":
    main()
