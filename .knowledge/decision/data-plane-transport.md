---
id: decision:data-plane-transport
type: decision
title: Data Plane Transport
---
How SQL connections from the test reach the template and each fork inside concept:server-process.

```yaml
decision:
  resolution: loopback TCP default; explicit Unix socket option implemented 2026-10-03 for requirement:go-external-test-backend; Go helper transport independent of embedded or subprocess execution (api:go-test-fixture)
  options:
    tcp_per_server:
      pros: [reuses the Go Server as is since api:clone already returns a listening server, every driver on every OS works (system:postgres-drivers), one pool per fork is natural]
      cons: [one ephemeral port per live fork; bounded by policy:fork-pool-limit so harmless]
    uds_per_server:
      shape: -socket-dir DIR flag; each server creates a private child directory with .s.PGSQL.5432; endpoint host is that directory, port is the socket identifier, DSN includes host query parameter
      support: Go unix means SOCK_STREAM; modern Windows supports this; unixgram and unixpacket are unsupported there; https://go.dev/src/net/unixsock.go
      pros: [no TCP port allocation, ordinary PostgreSQL Unix socket drivers, latency can be measured separately]
      cons: [pgjdbc needs junixsocket, driver support varies on Windows; OS AF_UNIX stream is supported, stale socket files, ~104 byte path limit]
    single_endpoint_routed:
      shape: one listener; the startup packet database name selects template or fork inside pgmem.go startSession
      pros: [one port and one URL prefix, a fork only changes dbname]
      cons: [new routing layer in pgmem.go, fork ids leak into dbname, pools are still per fork]
  wrapper_view: the wrapper hands out a DSN or JDBC URL; transport is invisible above that line
  next: other language Unix socket convenience options may follow; control channel stays independent from SQL transport
  later: routed endpoint only if port usage ever hurts
```
