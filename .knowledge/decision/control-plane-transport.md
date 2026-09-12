---
id: decision:control-plane-transport
type: decision
title: Control Plane Transport
---
Which channel carries snapshot, fork and close commands between the wrapper and concept:server-process.

```yaml
decision:
  resolution: stdio JSON lines, api:control-protocol (implemented 2026-09-12)
  options:
    stdio:
      pros: [no socket to name, secure or clean up, same on Windows and unix, pipe EOF doubles as parent-death detection (decision:fork-release-detection), it is already the ready-line channel]
      cons: [one client per process, stdout must stay free of stray prints, wrapper needs a reader thread]
    control_socket:
      shape: uds or loopback TCP listener plus a token, address printed in the ready line
      pros: [several clients (xdist workers, forked JVMs) could share one process and one prepared template]
      cons: [needs auth and file cleanup, uds path length and Windows differences, a listener has no peer so parent death is not observable]
    sql_embedded:
      shape: wrapper runs a magic statement on a template connection and pgmem.go intercepts it
      pros: [no second channel]
      cons: [intercepting queries in the proxy is fragile, fork endpoint must be smuggled through a result set, close needs a live connection]
  later: a control socket can be added for shared-process mode without changing the message format
```
