# Version Map

Stable pgmem versions use `1.X.Y`; beta versions use `0.X.Y`. `X` identifies the bundled PostgreSQL major version. `Y` counts pgmem releases for that major and track. Every release artifact uses the same full `X.Y.Z` version.

| pgmem version | Bundled PostgreSQL | Track | Meaning |
|---|---|---|---|
| `1.18.x` | PostgreSQL `18.x` (`1.18.0` uses `18.3`) | Stable | Current stable line; `Y` advances for each pgmem release. |
| `0.19.x` | PostgreSQL 19 prereleases (initially `Beta 4`) | Beta | The initial build, `0.19.0`, uses `Beta 4`. |
| `1.19.x` | PostgreSQL `19.x` after general availability | Stable | Stable PostgreSQL 19 line, starting at `1.19.0`. |

## バージョン対応表

安定版は `1.X.Y`、ベータ版は `0.X.Y` です。`X` は同梱する PostgreSQL のメジャーバージョン、`Y` はそのメジャー系列とトラックでの pgmem のリリース番号です。すべての成果物に同じ完全な `X.Y.Z` を付けます。

| pgmem のバージョン | 同梱する PostgreSQL | トラック | 対応 |
|---|---|---|---|
| `1.18.x` | PostgreSQL `18.x`（`1.18.0` は `18.3`） | 安定版 | 現行の安定系列です。pgmem のリリースごとに `Y` を増やします。 |
| `0.19.x` | PostgreSQL 19 のベータ版（初回は `Beta 4`） | ベータ版 | 初回ビルド `0.19.0` では `Beta 4` を使います。 |
| `1.19.x` | PostgreSQL `19.x`（正式リリース後） | 安定版 | PostgreSQL 19 の安定系列で、`1.19.0` から始めます。 |
