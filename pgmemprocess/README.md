# Go subprocess tests / Go の外部プロセステスト

## English

`pgmemprocess` owns the pgmem binary and exposes TCP or Unix socket endpoints.
It and the shared `pgmemfixture` driver helpers do not import the embedded engine.
Install the binary once, save the test below in an existing Go module, and run it:

```sh
go install github.com/shibukawa/pgmem/cmd/pgmem@latest
go get github.com/shibukawa/pgmem/pgmemprocess
go test ./...
```

```go title="store_test.go"
package store_test

import (
    "context"
    "database/sql"
    "os"
    "testing"

    "github.com/shibukawa/pgmem/pgmemprocess"
)

var fx *pgmemprocess.Fixture

func TestMain(m *testing.M) {
    os.Exit(pgmemprocess.Run(m, pgmemprocess.FixtureOptions{
        Options: pgmemprocess.Options{Transport: "tcp", Database: "app"},
        Prepare: func(ctx context.Context, db *sql.DB, _ string) error {
            _, err := db.ExecContext(ctx, "CREATE TABLE orders(id int)")
            return err
        },
    }, func(f *pgmemprocess.Fixture) { fx = f }))
}

func TestCreateOrder(t *testing.T) {
    t.Parallel()
    testDB := fx.For(t)
    db := testDB.DB()
    if _, err := db.Exec("INSERT INTO orders VALUES (1)"); err != nil {
        t.Fatal(err)
    }
    // testDB.DSN() and testDB.PgxPool() use this same database.
}
```

Binary lookup is `Options.Binary`, `PGMEM_BINARY`, then `pgmem` on `PATH`.
There is no runtime download. Set `Transport: "unix"` on hosts supporting AF_UNIX stream sockets, including modern Windows;
`SocketDir` is an existing short parent directory, default `/tmp` on Unix or `os.TempDir()` on Windows. Each server
has its own private socket directory. Normal shutdown removes it, and reset
keeps the endpoint and client sockets. Unsupported hosts fail at listen time. This mode uses stream sockets, not datagrams. Database data
stays in memory; socket entries are the only host filesystem artifacts. A forced
kill can leave socket paths; remove only the abandoned private directory.

`fx.For(t)` acquires one fork and owns teardown. Its `DB()`, `PgxConn()`,
`PgxPool()` and `DSN()` use the same database. `Reset(ctx)` restores it in place.
`SharedDB()` and `SharedPgxPool()` on the fixture share a read-only-by-convention
fork and close with the fixture. Acquire shared handles during setup when possible.
Each shared fork occupies one pool slot; do not write through shared handles.

Startup, snapshot and fixture fork waits default to 30 seconds. Shutdown is
bounded by 10 seconds. Override the corresponding option to change a deadline.
Close active transactions before snapshot or reset. `Start`, `Server.Snapshot`,
`Snapshot.Fork` and `Server.Restore` are available for manual lifecycles.

## 日本語

`pgmemprocess` は pgmem バイナリを所有し、TCP または Unix socket の接続先を渡します。
共通のドライバ用 helper `pgmemfixture` とともに、組み込みエンジンを import しません。
バイナリを一度インストールし、既存の Go module に次のテストを保存して実行します。

```sh
go install github.com/shibukawa/pgmem/cmd/pgmem@latest
go get github.com/shibukawa/pgmem/pgmemprocess
go test ./...
```

```go title="store_test.go"
package store_test

import (
    "context"
    "database/sql"
    "os"
    "testing"

    "github.com/shibukawa/pgmem/pgmemprocess"
)

var fx *pgmemprocess.Fixture

func TestMain(m *testing.M) {
    os.Exit(pgmemprocess.Run(m, pgmemprocess.FixtureOptions{
        Options: pgmemprocess.Options{Transport: "tcp", Database: "app"},
        Prepare: func(ctx context.Context, db *sql.DB, _ string) error {
            _, err := db.ExecContext(ctx, "CREATE TABLE orders(id int)")
            return err
        },
    }, func(f *pgmemprocess.Fixture) { fx = f }))
}

func TestCreateOrder(t *testing.T) {
    t.Parallel()
    testDB := fx.For(t)
    db := testDB.DB()
    if _, err := db.Exec("INSERT INTO orders VALUES (1)"); err != nil {
        t.Fatal(err)
    }
    // testDB.DSN() と testDB.PgxPool() も同じデータベースを使います。
}
```


`Options.Binary`、`PGMEM_BINARY`、`PATH` 上の `pgmem`
の順でバイナリを探し、実行時にダウンロードしません。AF_UNIX の stream socket に対応するホスト（対応する Windows も含む）では
`Transport: "unix"` を指定できます。`SocketDir` は既存の短い親ディレクトリで、
既定は Unix で `/tmp`、Windows で `os.TempDir()` です。各サーバーが専用の socket ディレクトリを持ち、正常終了時に削除します。
reset では接続先とクライアントの socket を維持します。非対応ホストでは listen 時にエラーになります。このモードは datagram ではなく stream socket を使います。
データベースのデータはメモリに保持し、ホストのファイルシステムに置くのは socket のエントリだけです。
強制終了で残った場合は、そのプロセスが使っていた専用ディレクトリだけを削除します。

`fx.For(t)` はフォークを 1 つ取得し、後片付けを管理します。`DB()`、`PgxConn()`、
`PgxPool()`、`DSN()` はすべて同じデータベースを使います。`Reset(ctx)` は接続先を変えずに戻します。
fixture の `SharedDB()` と `SharedPgxPool()` は読み取り専用として使うフォークを共有し、
fixture とともに閉じます。できればセットアップ時に取得してください。共有フォークもプールの
1 枠を占有します。共有 handle には書き込まないでください。

起動、snapshot、fixture の fork 待ちは既定で 30 秒、終了待ちは 10 秒です。
対応する option で期限を変更できます。snapshot や reset の前には、開いたトランザクションを
終了させます。手動管理には `Start`、`Server.Snapshot`、`Snapshot.Fork`、`Server.Restore` が使えます。
