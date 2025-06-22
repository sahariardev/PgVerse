package core

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"testing"
	"time"
)

func TestProxyServer(t *testing.T) {

	go func() {
		err := StartServer("5433", RequestHandler)
		if err != nil {
			return
		}
	}()

	ctx := context.Background()
	url := "postgres://testUser:testPass@localhost:5433/testDB?sslmode=disable"

	pool, err := pgxpool.New(ctx, url)

	if err != nil {
		t.Error(err)
	}

	defer pool.Close()

	var now time.Time
	err = pool.QueryRow(ctx, "SELECT NOW()").Scan(&now)

	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}

	fmt.Printf("Current time from PostgreSQL: %v", now)

	fmt.Println("Success")

}
