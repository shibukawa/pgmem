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

選定基準は単純です。主要なマネージドサービス（Amazon RDS、Google Cloud SQL）と PGlite の多くが提供している拡張を、ORM から使われるものを優先して入れます。リリースに入っていない拡張を入れた独自ビルドの pgmem も、同じ手順で作れます。

### エージェントに任せる

リポジトリには、コーディングエージェントにこの手順を一通り実行させるスキル `add-pgmem-extension` が入っています。pgmem のチェックアウト内では Claude Code が `/add-pgmem-extension` として自動で見つけます。他のエージェントやフォークで使うときは、リポジトリからインストールします。

```bash
npx skills add shibukawa/pgmem --skill add-pgmem-extension
```

あとは contrib モジュールの名前を挙げて、普通の言葉で頼みます。

> pgmem に pg_surgery を追加して。

スキルはエージェントにドライバースクリプト `skills/add-pgmem-extension/add-extension.sh` を示し、エージェントは次のことをします。

1. `add-extension.sh status` で、同梱済みの拡張と、上流に回帰テストが残っている未同梱の contrib モジュールを一覧します。ピン留めされた PostgreSQL のソースが `wasm/out/src` になければダウンロードします。
2. `add-extension.sh add <name>` で、`wasm/build.sh` の `CONTRIB_MODULES` に名前を追加し、`contrib/<name>/sql`、`expected`、`data` を `testdata/regress/<name>` にコピーし、Makefile の `REGRESS` の並びを `regress_test.go` に登録します。回帰テストランナーが再現できない psql コマンド、写す必要のある上流の `REGRESS_OPTS`、モジュールがリンクするホストのライブラリがあれば報告します。
3. エージェントが拡張の一覧を手で編集します。`README.md`、`.knowledge/policy/bundled-extensions.md`、そしてこのページの両言語版です。
4. `add-extension.sh build` で `wasm/build.sh`（Devbox 経由の Emscripten で `postgres.wasm` と `internal/assets/share.tar.gz` を生成）と `wasm/gen-aot.sh`（wasm2go のフォークで `internal/aot/pgaot` を再生成）を実行します。数分かかり、wasm ツールチェーンが必要なのはこのステップだけです。
5. `add-extension.sh smoke <name>` で pgmem を起動して `CREATE EXTENSION` を実行し、インストールされたオブジェクトを表示します。`add-extension.sh test <name>` で取り込んだ回帰テストを再生し、上流の期待出力と突き合わせます。
6. 1 拡張につき 1 コミット、`pgmem: bundle <name>` でコミットします。

PostgreSQL のソースツリー外の拡張は pgvector の形式に従います。リポジトリ・バージョン・チェックサムを書いたロックファイルと、`build.sh` 内の専用のコンパイル手順です。スキルはこの形式を説明しますが、ドライバーは自動化していません。

### 手作業で

ドライバーを使わない場合の同じ手順です。

1. `wasm/build.sh` の `CONTRIB_MODULES` に contrib の名前を追加します。control ファイル、SQL、データファイルは、モジュールと一緒に配布される share ツリーにコピーされます。
2. PostgreSQL のソースツリー外の拡張は、pgvector と同じように、リポジトリ・バージョン・チェックサムを書いたロックファイルと、`build.sh` のコンパイル手順を追加します。
3. `./wasm/build.sh` でリビルドし、`./wasm/gen-aot.sh` で Go を再生成します。
4. 拡張の上流の回帰テストファイルを `testdata/regress/<name>` に取り込み、`regress_test.go` に並べます。テストスイートがそれを pgmem に対して再生します。
5. 1 コミットにつき 1 拡張でコミットします。

各モジュールは `_PG_init` とマジックブロックのシンボル名を変えてコンパイルし、生成したテーブルを通して、静的な `dlopen` の代替実装がそれを見つけます。動的リンクなしで `CREATE EXTENSION` が動くのはこの仕組みのおかげです。
