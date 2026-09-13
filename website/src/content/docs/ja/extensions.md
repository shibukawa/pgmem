---
title: 拡張機能
description: pgmem に同梱している PostgreSQL の拡張と、新しい拡張の追加手順。
---

拡張は pgmem のビルドに静的にリンクしています。下の表にある名前は、何もインストールせずに `CREATE EXTENSION` できます。

## 同梱している拡張

| 分類 | 拡張 |
|---|---|
| 手続き言語 | `plpgsql` |
| 型と演算子 | `citext`、`hstore`、`ltree`、`cube`、`seg`、`isn`、`intarray`、`lo`、`earthdistance` |
| 全文検索とあいまい一致 | `pg_trgm`、`fuzzystrmatch`、`unaccent`、`dict_int`、`dict_xsyn` |
| インデックス | `btree_gist`、`btree_gin`、`bloom` |
| ベクトル | `vector`（pgvector 0.8.6） |
| 暗号と ID | `pgcrypto`、`uuid-ossp` |
| テーブル関数とサンプリング | `tablefunc`、`tsm_system_rows`、`tsm_system_time` |
| 調査と統計 | `pg_stat_statements`、`pgstattuple`、`pageinspect`、`pg_buffercache`、`pg_freespacemap`、`pg_visibility`、`amcheck`、`pg_prewarm` |
| ロード可能モジュール | `auto_explain`（`CREATE EXTENSION` ではなく `LOAD 'auto_explain'`） |

Snowball のステミング辞書と、すべてのエンコーディング変換もリンク済みです。

**OpenSSL なしの pgcrypto。** pgcrypto の OpenSSL 層と zlib 層は、Go の `crypto`、`math/big`、`compress` パッケージの呼び出しに置き換えています。ダイジェスト、AES、Blowfish、DES、3DES、CAST5、RSA や ElGamal 鍵による OpenPGP 暗号化、OpenPGP の圧縮まで動き、GnuPG で作ったメッセージも復号できます。`fips_mode()` は常に false で、bzip2 圧縮の OpenPGP パケットは上流と同じく未対応です。

## 同梱していない拡張

表にない拡張は `module is not linked into this pgmem build` というエラーになります。次のものは設計上の対象外です。

- `pg_cron` や `pg_partman` のワーカーのようなバックグラウンドワーカー。シングルユーザーモードには起動役のポストマスターがいないためです。
- `dblink` や `postgres_fdw` のような外部ネットワークへの接続。
- `plv8`、`plperl`、`plpython` のような別言語のランタイム。
- PostGIS。サイズと依存関係の重さのためです。

これらが必要なテストにはコンテナを使ってください。

## 拡張の追加手順

選定基準は単純です。主要なマネージドサービス（Amazon RDS、Google Cloud SQL）と PGlite の多くが提供している拡張を、ORM から使われるものを優先して入れます。

1. `wasm/build.sh` の `CONTRIB_MODULES` に contrib の名前を追加します。control ファイル、SQL、データファイルは、モジュールと一緒に配布される share ツリーにコピーされます。
2. PostgreSQL のソースツリー外の拡張は、pgvector と同じように、リポジトリ・バージョン・チェックサムを書いたロックファイルと、`build.sh` のコンパイル手順を追加します。
3. `./wasm/build.sh` でリビルドし、`./wasm/gen-aot.sh` で Go を再生成します。
4. 拡張の上流の回帰テストファイルを `testdata/regress/<name>` に取り込み、`regress_test.go` に並べます。テストスイートがそれを pgmem に対して再生します。
5. 1 コミットにつき 1 拡張でコミットします。

各モジュールは `_PG_init` とマジックブロックのシンボル名を変えてコンパイルし、生成したテーブルを通して、静的な `dlopen` の代替実装がそれを見つけます。動的リンクなしで `CREATE EXTENSION` が動くのはこの仕組みのおかげです。
