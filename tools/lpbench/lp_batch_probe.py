#!/usr/bin/env python3
"""Measure the log poller's current unfinalized-block walk vs a batched walk.

This is deliberately a transport-level probe: it issues exactly the same
eth_getBlockByNumber / eth_getLogs calls as the production log poller, but
without the database, head tracker or reorg handler.  It answers the question
"would batching help on this RPC endpoint, and at what chunk size?" before we
change chainlink-evm.

Usage:
  # compare two real endpoints over 300 finalized Arbitrum blocks
  python3 lp_batch_probe.py \
      --rpc alchemy=https://... --rpc linkpool=https://... \
      --blocks 300 --batch-sizes 1,2,4,8,16,32 --runs 3 \
      --address 0x... --topic0 0x...

  # local validation against the mock server
  python3 mock_rpc_server.py --port 8545 --delay-ms 160 --batch-concurrency 16 &
  python3 lp_batch_probe.py --rpc mock=http://127.0.0.1:8545 \
      --blocks 32 --batch-sizes 1,2,4,8,16 --runs 1

The script never prints RPC URLs, only the labels you give them.
"""

import argparse
import http.client
import json
import random
import re
import statistics
import ssl
import sys
import time
import urllib.parse


class RPCError(RuntimeError):
    pass


class RPC:
    """Minimal keep-alive JSON-RPC client. One instance == one connection."""

    def __init__(self, url: str, timeout: float = 60.0):
        self.url = url
        u = urllib.parse.urlparse(url)
        if u.scheme not in ("http", "https"):
            raise ValueError(f"unsupported scheme: {u.scheme}")
        self.https = u.scheme == "https"
        self.host = u.hostname
        self.port = u.port or (443 if self.https else 80)
        self.path = u.path or "/"
        if u.query:
            self.path += "?" + u.query
        self.timeout = timeout
        self.conn = None
        self.next_id = 0
        self.bytes_sent = 0
        self.bytes_recv = 0

    def _connect(self):
        if self.https:
            return http.client.HTTPSConnection(
                self.host, self.port, timeout=self.timeout,
                context=ssl.create_default_context(),
            )
        return http.client.HTTPConnection(self.host, self.port, timeout=self.timeout)

    def _close(self):
        if self.conn is not None:
            try:
                self.conn.close()
            except Exception:
                pass
            self.conn = None

    def _post(self, payload):
        body = json.dumps(payload).encode()
        self.bytes_sent += len(body)
        last_err = None
        for attempt in range(2):
            try:
                if self.conn is None:
                    self.conn = self._connect()
                self.conn.request(
                    "POST", self.path, body=body,
                    headers={"Content-Type": "application/json",
                             "Accept": "application/json",
                             "Connection": "keep-alive"},
                )
                resp = self.conn.getresponse()
                raw = resp.read()
                self.bytes_recv += len(raw)
                if resp.status != 200:
                    raise RPCError(f"HTTP {resp.status}: {raw[:300]!r}")
                return json.loads(raw)
            except RPCError:
                # HTTP 4xx (batch too large, 429, auth) is a real answer from
                # the provider, not a broken connection.  Do not retry it.
                self._close()
                raise
            except (http.client.HTTPException, OSError) as e:
                last_err = e
                self._close()
                if attempt == 0:
                    continue
                raise
        raise last_err  # pragma: no cover

    def batch(self, calls):
        """calls = [(method, params), ...] -> (results, elapsed_seconds)."""
        if not calls:
            return [], 0.0
        ids = []
        payload = []
        for method, params in calls:
            self.next_id += 1
            ids.append(self.next_id)
            payload.append({"jsonrpc": "2.0", "id": self.next_id,
                            "method": method, "params": params})

        start = time.perf_counter()
        data = self._post(payload)
        elapsed = time.perf_counter() - start

        if isinstance(data, dict):
            data = [data]
        by_id = {item.get("id"): item for item in data if isinstance(item, dict)}
        results = []
        for i in ids:
            item = by_id.get(i)
            if item is None:
                raise RPCError(f"missing response for batch id {i}")
            if item.get("error"):
                raise RPCError(f"rpc error: {item['error']}")
            results.append(item.get("result"))
        return results, elapsed

    def call(self, method, params):
        results, elapsed = self.batch([(method, params)])
        return results[0], elapsed


def chunks(seq, n):
    for i in range(0, len(seq), n):
        yield seq[i:i + n]


def get_block(rpc: RPC, number: int):
    result, elapsed = rpc.call("eth_getBlockByNumber", [hex(number), False])
    if result is None:
        raise RPCError(f"block {number} not available")
    return result, elapsed


def log_filter(block_hash, addresses, topics):
    f = {"blockHash": block_hash}
    if addresses:
        f["address"] = addresses
    if topics:
        f["topics"] = [topics]  # topic0 position only
    return f


def run_legacy(rpc: RPC, start: int, end: int, addresses, topics):
    """Exact shape of today's getUnfinalizedLogs: N log calls + N header calls."""
    t0 = time.perf_counter()
    anchor, _ = get_block(rpc, start)
    current = anchor
    logs = []
    header_calls = 1
    log_calls = 0

    res, _ = rpc.call("eth_getLogs", [log_filter(current["hash"], addresses, topics)])
    logs.extend(res)
    log_calls += 1

    for n in range(start + 1, end + 1):
        nxt, _ = rpc.call("eth_getBlockByNumber", [hex(n), False])
        header_calls += 1
        if nxt is None:
            raise RPCError(f"block {n} not available")
        if nxt["parentHash"] != current["hash"]:
            raise RPCError(f"reorg at {n}: parent {nxt['parentHash']} != {current['hash']}")
        res, _ = rpc.call("eth_getLogs", [log_filter(nxt["hash"], addresses, topics)])
        logs.extend(res)
        log_calls += 1
        current = nxt

    return {
        "elapsed": time.perf_counter() - t0,
        "header_calls": header_calls,
        "log_calls": log_calls,
        "logs": logs,
    }


def run_batched(rpc: RPC, start: int, end: int, addresses, topics, batch_size: int):
    """Fetch headers in batches, verify the parent chain, then logs by hash in batches."""
    t0 = time.perf_counter()
    anchor, _ = get_block(rpc, start)
    header_calls = 1
    log_calls = 0

    numbers = list(range(start + 1, end + 1))
    headers = []
    for chunk in chunks(numbers, batch_size):
        reqs = [("eth_getBlockByNumber", [hex(n), False]) for n in chunk]
        results, _ = rpc.batch(reqs)
        header_calls += len(reqs)
        for n, h in zip(chunk, results):
            if h is None:
                raise RPCError(f"block {n} not available")
            headers.append(h)

    prev = anchor
    for h in headers:
        if h["parentHash"] != prev["hash"]:
            raise RPCError(f"reorg at {h['number']}: parent {h['parentHash']} != {prev['hash']}")
        prev = h

    blocks = [anchor] + headers
    logs = []
    for chunk in chunks(blocks, batch_size):
        reqs = [("eth_getLogs", [log_filter(b["hash"], addresses, topics)]) for b in chunk]
        results, _ = rpc.batch(reqs)
        log_calls += len(reqs)
        for res in results:
            logs.extend(res)

    return {
        "elapsed": time.perf_counter() - t0,
        "header_calls": header_calls,
        "log_calls": log_calls,
        "logs": logs,
    }


def normalize_logs(logs):
    out = []
    for l in logs:
        out.append((
            l.get("blockHash", "").lower(),
            l.get("transactionHash", "").lower(),
            int(l.get("logIndex", "0x0"), 16),
            l.get("address", "").lower(),
            tuple(t.lower() for t in l.get("topics", [])),
            l.get("data", "").lower(),
        ))
    return sorted(out)


def fmt_ms(seconds):
    return f"{seconds * 1000:,.0f}"


def main():
    p = argparse.ArgumentParser()
    p.add_argument("--rpc", action="append", required=True, metavar="LABEL=URL",
                   help="endpoint to test; repeat for multiple endpoints")
    p.add_argument("--blocks", type=int, default=300,
                   help="number of consecutive blocks to walk (default 300)")
    p.add_argument("--batch-sizes", default="1,2,4,8,16,32,50",
                   help="comma-separated chunk sizes (default 1,2,4,8,16,32,50)")
    p.add_argument("--runs", type=int, default=3,
                   help="runs per batch size; median is reported (default 3)")
    p.add_argument("--latest-offset", type=int, default=10_000,
                   help="use blocks this far behind the reference head (default 10000)")
    p.add_argument("--from-block", type=int, default=None,
                   help="override start block (must be finalized)")
    p.add_argument("--to-block", type=int, default=None,
                   help="override end block (must be finalized)")
    p.add_argument("--address", action="append", default=[],
                   help="address filter; repeat for multiple")
    p.add_argument("--topic0", action="append", default=[],
                   help="topic0 filter; repeat for multiple")
    p.add_argument("--timeout", type=float, default=120.0)
    p.add_argument("--seed", type=int, default=1,
                   help="shuffle batch-size order to spread provider cache effects (default 1)")
    args = p.parse_args()

    endpoints = []
    for spec in args.rpc:
        if "=" not in spec:
            p.error(f"--rpc must be LABEL=URL, got {spec!r}")
        label, url = spec.split("=", 1)
        endpoints.append((label, url))

    batch_sizes = [int(x) for x in args.batch_sizes.split(",") if x.strip()]
    addresses = [a.lower() for a in args.address]
    topics = [t.lower() for t in args.topic0]
    for a in addresses:
        if not re.fullmatch(r"0x[0-9a-f]{40}", a):
            p.error(f"--address must be 0x followed by 40 hex characters, got {a!r}")
    for t in topics:
        if not re.fullmatch(r"0x[0-9a-f]{64}", t):
            p.error(f"--topic0 must be 0x followed by 64 hex characters, got {t!r}")
    if not addresses and not topics:
        print("WARNING: no --address/--topic0 filter; eth_getLogs will return every log "
              "in each block, which is not production-like and may hit provider limits.\n",
              flush=True)

    # Resolve the range once, behind the lowest head, so every endpoint has it.
    heads = {}
    for label, url in endpoints:
        try:
            r = RPC(url, timeout=args.timeout)
            head_hex, _ = r.call("eth_blockNumber", [])
            heads[label] = int(head_hex, 16)
        except Exception as e:
            print(f"WARNING: cannot read head from {label}: {e}", flush=True)
    if not heads:
        p.error("no endpoint returned eth_blockNumber")
    head = min(heads.values())
    to_block = args.to_block if args.to_block is not None else head - args.latest_offset
    from_block = args.from_block if args.from_block is not None else to_block - args.blocks + 1
    if from_block < 0 or to_block <= from_block:
        p.error(f"bad range {from_block}..{to_block} (head={head})")
    n_blocks = to_block - from_block + 1
    print(f"range {from_block}..{to_block} ({n_blocks} blocks), head={head}\n", flush=True)

    rows = []
    for label, url in endpoints:
        try:
            rpc = RPC(url, timeout=args.timeout)
        except ValueError as e:
            print(f"== {label} ==\n  SKIP: {e}\n", flush=True)
            continue
        print(f"== {label} ==", flush=True)

        # Warm up the connection and prove the range is available.
        try:
            get_block(rpc, from_block)
            get_block(rpc, to_block)
        except Exception as e:
            print(f"  SKIP: {e}\n", flush=True)
            continue

        print("  legacy baseline...", end="", flush=True)
        try:
            legacy = run_legacy(rpc, from_block, to_block, addresses, topics)
        except Exception as e:
            print(f" SKIP: legacy baseline failed: {e}\n", flush=True)
            continue
        legacy_bps = n_blocks / legacy["elapsed"]
        avg_call_ms = legacy["elapsed"] * 1000 / max(1, legacy["header_calls"] + legacy["log_calls"])
        print(f" {legacy_bps:,.1f} blocks/s "
              f"({fmt_ms(legacy['elapsed'])} ms, {legacy['header_calls']} hdr + "
              f"{legacy['log_calls']} log calls, ~{avg_call_ms:,.0f} ms/call)", flush=True)
        ref_logs = normalize_logs(legacy["logs"])

        # Shuffle so a provider-side cache from one batch size does not
        # systematically favour the next one.  Results are sorted below.
        order = batch_sizes[:]
        random.Random(args.seed).shuffle(order)
        for bs in order:
            times = []
            logs_ok = True
            err = None
            for _ in range(max(1, args.runs)):
                try:
                    res = run_batched(rpc, from_block, to_block, addresses, topics, bs)
                    times.append(res["elapsed"])
                    if normalize_logs(res["logs"]) != ref_logs:
                        logs_ok = False
                except Exception as e:
                    err = str(e)
                    break
            if err:
                print(f"  batch {bs:>3}: FAILED: {err}", flush=True)
                continue
            median = statistics.median(times)
            bps = n_blocks / median
            speedup = bps / legacy_bps if legacy_bps else 0
            rows.append((label, bs, median, bps, speedup, logs_ok))
            print(f"  batch {bs:>3}: {bps:>8,.1f} blocks/s "
                  f"({fmt_ms(median):>7} ms, {speedup:>5.1f}x vs legacy)"
                  f"{'' if logs_ok else '  *** LOG MISMATCH ***'}", flush=True)
        print(flush=True)

    print("SUMMARY")
    print(f"{'endpoint':<20} {'batch':>6} {'ms':>10} {'blocks/s':>10} {'speedup':>8} {'logs':>6}")
    for label, bs, median, bps, speedup, logs_ok in sorted(rows, key=lambda r: (r[0], r[1])):
        print(f"{label:<20} {bs:>6} {median*1000:>10,.0f} {bps:>10,.1f} {speedup:>8.1f}x "
              f"{'ok' if logs_ok else 'MISMATCH':>6}")


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        sys.exit(130)
