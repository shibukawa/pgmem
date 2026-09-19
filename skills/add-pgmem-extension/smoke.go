// Command smoke starts pgmem, creates one extension and prints what it
// installed. It is the check that a freshly bundled module is linked and
// its control/SQL files are in the share tree:
//
//	go run ./skills/add-pgmem-extension pg_trgm
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/shibukawa/pgmem"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: smoke <extension>")
		os.Exit(2)
	}
	name := os.Args[1]
	ctx := context.Background()
	s, err := pgmem.Start(ctx, pgmem.Options{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "start:", err)
		os.Exit(1)
	}
	defer s.Close()
	conn, err := pgx.Connect(ctx, s.DSN())
	if err != nil {
		fmt.Fprintln(os.Stderr, "connect:", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	if _, err := conn.Exec(ctx, fmt.Sprintf("CREATE EXTENSION %s", pgx.Identifier{name}.Sanitize())); err != nil {
		fmt.Fprintln(os.Stderr, "CREATE EXTENSION failed:", err)
		os.Exit(1)
	}
	var version string
	if err := conn.QueryRow(ctx, "SELECT extversion FROM pg_extension WHERE extname = $1", name).Scan(&version); err != nil {
		fmt.Fprintln(os.Stderr, "pg_extension:", err)
		os.Exit(1)
	}
	fmt.Printf("%s %s installed\n", name, version)
	rows, err := conn.Query(ctx, `SELECT pg_describe_object(classid, objid, objsubid)
		FROM pg_depend WHERE refclassid = 'pg_extension'::regclass
		AND refobjid = (SELECT oid FROM pg_extension WHERE extname = $1) AND deptype = 'e'
		ORDER BY 1`, name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "pg_depend:", err)
		os.Exit(1)
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		var d string
		if err := rows.Scan(&d); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		n++
		if n <= 20 {
			fmt.Println("  " + d)
		}
	}
	if n > 20 {
		fmt.Printf("  ... %d objects in total\n", n)
	}
}
