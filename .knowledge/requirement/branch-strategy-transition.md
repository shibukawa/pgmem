---
id: requirement:branch-strategy-transition
type: requirement
title: Branch Strategy Transition
---
Separate the PostgreSQL 18 stable release line from newest-major development, and align repository automation with the branch roles (policy:branching, policy:versioning, flow:release).

```yaml
requirement:
  outcomes:
    - postgresql/18 contains pgmem 1.18.0 for PostgreSQL 18 and is the GitHub default branch until PostgreSQL 19 GA
    - main contains pgmem 0.19.0 for PostgreSQL 19 Beta 4 and remains the newest-major branch
    - after PostgreSQL 19 GA, main starts its stable release line at pgmem 1.19.0 and becomes the GitHub default branch
    - tag v1.18.0 points to the PostgreSQL 18 line; tag v0.19.0 points to main; future v1.19.0 points to main
    - CI validates pushes to both long-lived branches
    - documentation deploys only from the current GitHub default branch
  acceptance:
    - README documents branch roles, version/tag mapping, and the default-branch transition in English and Japanese
    - GitHub Actions run CI for main and postgresql/18
    - docs workflow builds both branches and deploys only the configured default branch
    - release workflow accepts vX.Y.Z and publishes only when the tag commit is reachable from its assigned branch
```
