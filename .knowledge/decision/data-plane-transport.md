---
id: decision:data-plane-transport
type: decision
title: Data Plane Transport
---
How SQL connections from the test reach the template and each fork inside concept:server-process.

```yaml
decision:
  resolution: loopback TCP, one listener per server or fork, port in the endpoint (implemented 2026-09-12); uds as a later opt-in
  options:
    tcp_per_server:
      pros: [reuses the Go Server as is since api:clone already returns a listening server, every driver on every OS works (system:postgres-drivers), one pool per fork is natural]
      cons: [one ephemeral port per live fork; bounded by policy:fork-pool-limit so harmless]
    uds_per_server:
      shape: -socket-dir DIR flag; endpoint carries socket DIR/.s.PGSQL.<n> instead of host and port
      pros: [no port allocation, lower latency than loopback, no macOS firewall prompt]
      cons: [pgjdbc needs junixsocket, Windows support uneven, stale socket files, ~104 byte path limit]
    single_endpoint_routed:
      shape: one listener; the startup packet database name selects template or fork inside pgmem.go startSession
      pros: [one port and one URL prefix, a fork only changes dbname]
      cons: [new routing layer in pgmem.go, fork ids leak into dbname, pools are still per fork]
  wrapper_view: the wrapper hands out a DSN or JDBC URL; transport is invisible above that line
  later: uds option for Python on unix once measured; routed endpoint only if port usage ever hurts
```
