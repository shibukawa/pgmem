---
id: requirement:multi-language-wrapper
type: requirement
title: Multi-Language Wrapper
---
Python and Java packages that bundle the pgmem binary so one dependency gives the same prepare-once, fork-per-test workflow as requirement:test-fixture-fork without Go.

```yaml
summary:
  scope_now: [api:python-wrapper, api:java-wrapper]  # both implemented 2026-09-12; publishing and the CI build matrix remain
  scope_later: Node.js; no standard DB API, so decide the driver story first
  layout:
    - packages/python  # PyPI package pgmem
    - packages/java    # Maven artifacts
  mechanism: concept:server-process spawned by the wrapper; api:control-protocol on stdio; SQL over decision:data-plane-transport
  lifecycle: flow:wrapper-test-lifecycle
  parity_with_go:
    - prepare once on a template, api:snapshot, api:clone per test
    - fork release guaranteed by the test framework lifecycle (decision:fork-release-detection)
    - parallel tests get independent forks; policy:fork-pool-limit applies
  distribution: policy:binary-distribution
  non_goals_now:
    - in-process embedding (JNI, ctypes); the process boundary is the design
    - replacing the client driver; users keep their normal one (system:postgres-drivers)
  precedent: zonky embedded-postgres (Java), ruff/uv platform wheels (Python), esbuild optionalDependencies (Node)
```
