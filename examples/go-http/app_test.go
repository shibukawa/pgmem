package orders

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/shibukawa/pgmem/pgmemprocess"
)

var fixture *pgmemprocess.Fixture

func TestMain(m *testing.M) {
	os.Exit(pgmemprocess.Run(m, pgmemprocess.FixtureOptions{
		Options: pgmemprocess.Options{Transport: "tcp", Database: "orders"},
		Prepare: func(ctx context.Context, db *sql.DB, _ string) error {
			_, err := db.ExecContext(ctx, "CREATE TABLE orders(id int PRIMARY KEY)")
			return err
		},
	}, func(f *pgmemprocess.Fixture) { fixture = f }))
}

func TestCreateOrder(t *testing.T) {
	t.Parallel()
	database := fixture.For(t)
	server := httptest.NewServer(NewHandler(database.DB()))
	t.Cleanup(server.Close) // HTTP server closes before the database handles.
	response, err := http.Post(server.URL+"/orders", "application/json", strings.NewReader(`{"id":42}`))
	if err != nil {
		t.Fatal(err)
	}
	response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		t.Fatalf("status: %d", response.StatusCode)
	}
	var id int
	if err := database.PgxPool().QueryRow(t.Context(), "SELECT id FROM orders").Scan(&id); err != nil || id != 42 {
		t.Fatalf("application write missing from same fork: %d %v", id, err)
	}
}

func TestAnotherTestStartsEmpty(t *testing.T) {
	t.Parallel()
	var count int
	if err := fixture.For(t).DB().QueryRow("SELECT count(*) FROM orders").Scan(&count); err != nil || count != 0 {
		t.Fatalf("another test leaked data: %d %v", count, err)
	}
}
