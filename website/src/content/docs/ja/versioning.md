---
title: バージョン番号
description: pgmem の安定版は 1.X.Y、ベータ版は 0.X.Y です。表で pgmem と同梱する PostgreSQL の対応を確認できます。
---

安定版は `1.X.Y`、ベータ版は `0.X.Y` です。`X` は同梱する PostgreSQL のメジャーバージョン、`Y` はそのメジャー系列とトラックでの pgmem のリリース番号です。すべての成果物に同じ完全な `X.Y.Z` を付けます。

| pgmem のバージョン | 同梱する PostgreSQL | トラック | 対応 |
|---|---|---|---|
| `1.18.x` | PostgreSQL `18.x`（`1.18.0` は `18.3`） | 安定版 | 現行の安定系列です。pgmem のリリースごとに `Y` を増やします。 |
| `0.19.x` | PostgreSQL 19 のベータ版（初回は `Beta 4`） | ベータ版 | 初回ビルド `0.19.0` では `Beta 4` を使います。 |
| `1.19.x` | PostgreSQL `19.x`（正式リリース後） | 安定版 | PostgreSQL 19 の安定系列で、`1.19.0` から始めます。 |

npm の dist-tag はパッケージごとに決めます。`0.X.Y` のベータ版には常に
`beta` を付けます。安定版の `latest` は、公開順ではなく、公開済み安定版のうち
SemVer が最も大きいバージョンを指します。現在の `latest` より古い安定版には、
同梱する PostgreSQL のメジャーバージョンが `X` となる `postgresql-X` を付けます。
ベータ版は `npm install @pgmem/core@beta`、PostgreSQL 18 の保守版系列は
`npm install @pgmem/core@postgresql-18` でインストールできます。
