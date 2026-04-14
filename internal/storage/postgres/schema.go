package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func InitSchema(db *sql.DB) error {
	const query = `
	CREATE TABLE IF NOT EXISTS tasks (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Error initializing schema: %w", err)
	}
	return nil
}
