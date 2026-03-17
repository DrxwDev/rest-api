package database

import (
	"context"
	"fmt"
	"time"

	"github.com/DrxwDev/rest-api/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(cfg config.DBConfig) (*pgxpool.Pool, error) {
	// 1. Parse configuration
	poolCfg, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	// 2. Fine-tune pool settings for high throughput
	poolCfg.MaxConns = 50
	poolCfg.MinConns = 10
	poolCfg.MaxConnLifetime = time.Hour
	poolCfg.MaxConnIdleTime = 30 * time.Minute
	poolCfg.HealthCheckPeriod = 30 * time.Second

	// 3. Create a timeout context specifically for the connection phase
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 4. Initialize the pool
	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	// 5. Verify the connection is actually alive
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancel()

	if err := pool.Ping(pingCtx); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return pool, nil
}
