---
title: 制限事項
description: ポストマスターが起動する通常の PostgreSQL と pgmem の違い。
---

## プロセスモデル

PostgreSQL は本番のサーバーと同じ形で動きます。postmaster が接続ごとにバックエンドプロセスを起動し、補助プロセス（checkpointer など）も動き、それらがメモリを共有します。各プロセスは独立した goroutine 上の WebAssembly インスタンスです。したがってセッションは本物です。接続同士のロックは待ち合い、デッドロックは検出され（`deadlock_timeout` 後に SQLSTATE 40P01）、`pg_stat_activity` と `pg_locks` にほかのセッションが見え、`SET`、一時テーブル、アドバイザリーロックといったセッション状態は接続ごとに独立しています。

- **`snapshot` と `reset` は開いているトランザクションを待ちます。** どちらもクラスタを停止して起動し直します。クライアントの接続はソケットを保ち、次のメッセージで新しいバックエンドに接続し直され、名前付きプリペアドステートメントと `LISTEN` の登録は作り直されます。先にトランザクションをコミットするか接続を閉じてください。ラッパーはタイムアウト後に `busy` エラーにします。
- **PostgreSQL 自身が終了させたバックエンド**（`pg_terminate_backend`、`DROP DATABASE ... WITH (FORCE)`、クラッシュ）は、サーバーと同じく FATAL メッセージとともに接続を終了します。
- **システムコールを一切行わない文のキャンセル**は、次にシステムコールが行われるまで遅れます。

## データベース

サーバーが持つすべてのデータベースに接続でき、`CREATE DATABASE` も動きます。存在しないデータベースは SQLSTATE 3D000 で拒否します。

## サーバーを閉じるとき

サーバーやフォークを閉じても、クライアントの接続は切りません。各接続は、クライアントが閉じるか、メッセージを送るか、30 秒たつまで開いたままです。メッセージには SQLSTATE 57P01 で応答します。node-postgres などのプールは、アイドル状態の接続を足元で切られると処理されないエラーを投げるためです。

## バックグラウンドプロセス

- 並列クエリと並列インデックス作成は動きます。ワーカーはクラスタのプロセスです。
- 独自のバックグラウンドワーカーを必要とする拡張は同梱していません。[拡張](../extensions/) を参照してください。
- `io_method` の既定は `sync` です（PostgreSQL 18 の `worker` は I/O ワーカープロセスを 3 つ増やすだけで利点がありません）。

## リソース

- 既定の `shared_buffers=32MB` で、サーバー 1 つあたりの常駐メモリはおよそ 150 MB です。生きているフォークはそれぞれ自分の複製を持ちます。大きなキャッシュが必要なテストでは、サーバーパラメータで `shared_buffers` を増やしてください。
- サーバーは 32 ビットの WebAssembly で、扱えるのは最大 4 GiB です。
- `statement_timeout` は、プロトコルメッセージの合間と、バックエンドが時計を読むかスリープするときに発火します。どちらもしない CPU ループの中では発火しません。

## フォーク数と待機

フォーク枠はスナップショットごとにあります。`MaxForks` はそのスナップショットから同時に起動できるフォーク数を制限し、テンプレートサーバーは枠に数えません。既定値はメモリから決まります。プロセスのメモリ上限（`GOMEMLIMIT`、cgroup の上限、物理メモリのうち最小のもの）の 4 分の 1 を、フォーク 1 個のコスト（`shared_buffers` + 約 32MB）で割った値で、7GB の CI ランナーと既定の `shared_buffers=32MB` なら 28 です。枠はフォークが実際に使うまでメモリを消費しません。メモリ量が取得できない環境では CPU 数になります。フォークごとにデータディレクトリとバッファキャッシュを複製するため、上限はメモリ使用量も抑えます。すべての枠が埋まると、次のフォーク要求は既存フォークが閉じて枠が空くまで待ちます。

| API | 上限の設定 | 空き枠を待つ時間の上限 |
|---|---|---|
| Go | `SnapshotOptions.MaxForks` または `pgmemtest.Options.MaxForks` | `Snapshot.Fork(ctx)` は context がキャンセルされるか期限を迎えるまで待ちます。サーバーロガーを有効にしている場合、5 秒を超える待機をログに出します。 |
| Python | `server.snapshot(max_forks=n)` | `snapshot.fork(timeout=秒)`。省略または `None` なら無期限に待ちます。期限を超えると `ProtocolError` の code は `pool_timeout` です。 |
| Java | `Server.snapshot(maxForks)` または JUnit の `.maxForks(n)` | `Snapshot.fork(Duration)` または JUnit の `.forkTimeout(Duration)`。既定ではタイムアウトしません。期限を超えると `pool_timeout` で失敗します。 |
| Node.js | `PgmemServer.start({ maxForks: n })` | `fork({ timeoutMs })` または `withFork(fn, { timeoutMs })`。省略すると無期限に待ちます。期限を超えると `PgmemError` の code は `pool_timeout` です。 |

JUnit 拡張で設定した上限は、登録したテンプレートごとに適用されます。スナップショットの `close()` は新しいフォークを拒否しますが、稼働中のフォークは止めません。Go の `Snapshot.Wait()` は、フォークがすべて閉じるまで待ちます。

## タイムアウトは待ちの種類ごとに異なる

フォーク枠を待つ時間は、スナップショット作成やサーバー起動を待つ時間を制限しません。それぞれ別の設定です。

| 待つもの | 設定 | 期限を超えたとき |
|---|---|---|
| スナップショット作成時に、開いているトランザクションが終わるのを待つ | Go: `Snapshot` の context。Python: `server.snapshot(timeout=30.0)`。Java: `Server.snapshot(maxForks, timeout)`（既定 30 秒）。Node.js: `snapshotTimeoutMs`（既定 30,000 ms）。 | `busy` で失敗します。テンプレートへの接続をコミットするか閉じてください。 |
| サーバープロセスの起動を待つ | Python: `pgmem.start(timeout=30)`。Java: `.readyTimeout(Duration)`（既定 30 秒）。Node.js: `startupTimeoutMs`（既定 30,000 ms）。Go: `Start` の context。 | 期限内にサーバーが ready にならなければ起動に失敗します。 |
| フォークをスナップショットへ戻す | Go: `Server.Reset(ctx)` または `Restore(ctx)`。Node.js: `fork.reset({ timeoutMs })`（既定 5,000 ms）。 | 開いたトランザクションが終わらなければ `busy` で失敗します。 |

## セキュリティ

TLS もパスワード認証もありません。サーバーは `127.0.0.1` だけで待ち受け、URL には `sslmode=disable` が付きます。ICU の照合順序は使えません。
