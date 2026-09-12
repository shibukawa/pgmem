---
id: rule:non-blocking-control-channel
type: rule
title: Non-Blocking Control Channel
---
Fork requests may block for a slot (policy:fork-pool-limit), so api:control-protocol responses are matched by id and may arrive out of order; a wrapper must never queue requests behind a pending fork.

```yaml
rule:
  deadlock: thread A waits in fork with a full pool; thread B's close, which would free the slot, sits behind A under a wrapper-side lock; both wait forever
  server: each request runs in its own goroutine; stdout writes are serialized by a mutex, one line per write
  wrapper: one reader thread dispatches responses to waiters by id (CompletableFuture, threading.Event); stdin writes lock only for the duration of one line
  ids: monotonic per wrapper process; echoed by the server; never reused while pending
  timeout: fork accepts timeout_ms and answers pool_timeout; wrappers surface it as a clear test failure instead of a hang
  shutdown: shutdown and stdin EOF fail every pending fork with server_exited
```
