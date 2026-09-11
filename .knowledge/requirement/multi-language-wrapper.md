---
id: requirement:multi-language-wrapper
type: requirement
title: Multi-Language Wrapper
---
Later phase: Python, Node.js, and Java packages that bundle a pgmem server binary so `import`/`require`/Maven dependency alone gives the same fixture-fork workflow.

```yaml
summary:
  mechanism: bundled platform binary spawned by the wrapper; control channel on stdin/stdout for snapshot and fork
  precedent: zonky embedded-postgres (Java), esbuild-style optionalDependencies (Node), platform wheels (Python)
  depends_on:
    - api:snapshot
    - api:clone
  not_now: true
```
