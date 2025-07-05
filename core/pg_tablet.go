package core

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Tablet struct {
	Config TabletConfig
	Pool   *pgxpool.Pool
}

type TabletConfig struct {
	ListenAddr string
	BackendDB  string
}

func NewTablet(config TabletConfig) (*Tablet, error) {
	pool, err := pgxpool.New(context.Background(), config.BackendDB)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to backend: %v", err)
	}

	return &Tablet{
		Config: config,
		Pool:   pool,
	}, nil
}

func (t *Tablet) ExecuteQuery(ctx context.Context, sql string) (pgx.Rows, error) {
	rows, err := t.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute query: %v", err)
	}
	return rows, nil
}
