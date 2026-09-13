// Program base opens a database/sql pool with pgx and nothing else; it is
// the baseline that withpgmem is compared against.
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	fmt.Println(db != nil, err)
}
