---
id: api:control-socket
type: api
title: Control Socket
---
Loopback listener serving api:control-protocol to many clients, so test workers can fork and reset against one concept:server-process (decision:node-test-integration).

```yaml
api:
  status: implemented 2026-09-13 (cmd/pgmem control.go listen and -control flag; socket_test.go)
  enable: pgmem -control 127.0.0.1:0; the ready event gains control {addr, token}
  env: PGMEM_CONTROL carries addr and token to workers
  framing: same JSON lines and ops as stdio; the first request on a connection must carry the token
  ownership: servers forked on a connection close when that connection ends, so a crashed worker cannot hold slots under policy:fork-pool-limit
  lifetime_signal: control-connection lifetime equals worker lifetime; decision:fork-release-detection rejected SQL connection counting, not this
  why_not_http: osmem-style HTTP has no connection to own forks
  stdio_owner: the spawning process keeps stdin; its EOF still ends everything
  later_use: e2e setups where the app server and the tests are different processes (Playwright webServer)
```
