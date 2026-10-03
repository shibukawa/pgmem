# Application tests / アプリケーションのテスト

## English

Run these examples from a checkout. Each application writes through its usual
client, and the test checks the write through the same fork. Go uses an owned
subprocess; Node keeps an import-time pool across sequential resets; Python
passes a selected engine into the application; Java injects one `TestDatabase`.

Build the binary once and export its absolute path:

```sh
go build -o /tmp/pgmem-example ./cmd/pgmem
export PGMEM_BINARY=/tmp/pgmem-example
```

On Windows, choose a local `.exe` output path and set `PGMEM_BINARY` to that
absolute path in your shell.

| Example | Run from the repository root | Source |
|---|---|---|
| Go HTTP handler | `go -C examples/go-http test -v` | [application](go-http/app.go), [tests](go-http/app_test.go) |
| Node HTTP server, Node 24+ | `npm --prefix examples/node-http install --omit=optional`, then `npm --prefix examples/node-http test` | [application](node-http/app.mjs), [setup](node-http/global-setup.mjs), [tests](node-http/app.test.mjs) |
| Python SQLAlchemy | `uv run --project examples/python-sqlalchemy pytest examples/python-sqlalchemy -q` | [application](python-sqlalchemy/app.py), [fixtures](python-sqlalchemy/conftest.py), [tests](python-sqlalchemy/test_app.py) |
| Java DataSource | `cd packages/java` then `./gradlew :pgmem-junit5:test --tests '*OrderApplicationTest'` | [complete example](../packages/java/pgmem-junit5/src/test/java/io/github/shibukawa/pgmem/junit5/examples/OrderApplicationTest.java) |

The Go module uses a local `replace`, Node uses a local package dependency,
and uv uses an editable local source. For use outside the checkout, install
the published pgmem packages and remove those local source settings. Clients,
HTTP servers and ORM engines are application-owned; close them before releasing
their fork. A reset restores database state. It does not reset the application's
memory, Valkey keys or OpenSearch indexes.

## 日本語

リポジトリの checkout から実行できます。各アプリが普段のクライアントで書き込み、
テストが同じフォークからその結果を検証します。Go は所有する外部プロセスを使い、
Node は import 時に作った pool を直列の reset 間で維持します。Python は選択済みの
engine をアプリに渡し、Java は一つの `TestDatabase` を注入します。

バイナリを一度ビルドし、絶対パスを環境変数に設定します。

```sh
go build -o /tmp/pgmem-example ./cmd/pgmem
export PGMEM_BINARY=/tmp/pgmem-example
```

Windows ではローカルの `.exe` 出力先を選び、その絶対パスを利用する shell で
`PGMEM_BINARY` に設定してください。

| サンプル | リポジトリのルートから実行するコマンド | ソース |
|---|---|---|
| Go HTTP ハンドラ | `go -C examples/go-http test -v` | [アプリ](go-http/app.go)、[テスト](go-http/app_test.go) |
| Node HTTP サーバー、Node 24 以降 | `npm --prefix examples/node-http install --omit=optional`、続けて `npm --prefix examples/node-http test` | [アプリ](node-http/app.mjs)、[準備](node-http/global-setup.mjs)、[テスト](node-http/app.test.mjs) |
| Python SQLAlchemy | `uv run --project examples/python-sqlalchemy pytest examples/python-sqlalchemy -q` | [アプリ](python-sqlalchemy/app.py)、[fixture](python-sqlalchemy/conftest.py)、[テスト](python-sqlalchemy/test_app.py) |
| Java DataSource | `cd packages/java` の後に `./gradlew :pgmem-junit5:test --tests '*OrderApplicationTest'` | [完全なサンプル](../packages/java/pgmem-junit5/src/test/java/io/github/shibukawa/pgmem/junit5/examples/OrderApplicationTest.java) |

Go module はローカルの `replace`、Node はローカルパッケージ依存、uv は編集可能な
ローカルソースを使います。checkout の外で使う場合は公開済みの pgmem パッケージを
導入し、これらのローカル設定を外します。クライアント、HTTP サーバー、ORM engine は
アプリ側で管理し、フォークを解放する前に閉じます。reset が戻すのは DB の状態です。
アプリのメモリ、Valkey のキー、OpenSearch の index は各サービス側で戻してください。
