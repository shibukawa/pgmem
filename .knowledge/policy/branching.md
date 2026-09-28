---
id: policy:branching
type: policy
title: Branching Strategy
---
Keep PostgreSQL major-version lines on separate branches while `main` tracks the newest major (policy:versioning, flow:release).

```yaml
policy:
  branches:
    main: newest PostgreSQL major; currently PostgreSQL 19 Beta 4 at pgmem 0.19.0; after PostgreSQL 19 GA, continue at stable 1.19.0
    postgresql/18: PostgreSQL 18 stable line, starting at pgmem 1.18.0
  default_branch:
    before_postgresql_19_ga: postgresql/18
    after_postgresql_19_ga: main
  release_tags:
    format: vX.Y.Z
    v1.18.x: postgresql/18
    v0.19.x: main
    v1.19.x: main after PostgreSQL 19 GA
  github_actions:
    ci_push_branches: [main, postgresql/18]
    docs_build_branches: [main, postgresql/18]
    docs_deploy: current GitHub default branch only
```
