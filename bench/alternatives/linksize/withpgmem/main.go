// Program withpgmem is base plus pgmem.Start; the size difference is what
// linking pgmem adds to a test binary.
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/shibukawa/pgmem"
)

func main() {
	s, err := pgmem.Start(context.Background(), pgmem.Options{})
	if err != nil {
		panic(err)
	}
	defer s.Close()
	db, err := sql.Open("pgx", s.DSN())
	fmt.Println(db != nil, err)
}
