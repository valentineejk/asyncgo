package database

import (
	"asyncgo/config"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func NewPostgres(cfg config.Config) (*sql.DB, error) {
	dsn := cfg.DatabaseUrl()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database")
	}
	return db, nil
}
