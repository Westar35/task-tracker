package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func InitSchema(db *sql.DB) error {
	const query2 = `
	CREATE TABLE IF NOT EXISTS users (
		id BIGSERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`

	ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel2()

	_, err := db.ExecContext(ctx2, query2)
	if err != nil {
		return fmt.Errorf("Error initializing users table: %w", err)
	}

	const query1 = `
	CREATE TABLE IF NOT EXISTS tasks (
		id BIGSERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		status TEXT NOT NULL,
		user_id BIGINT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		CONSTRAINT fk_tasks_user
			FOREIGN KEY(user_id) REFERENCES users(id)
			ON DELETE CASCADE
	);
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err = db.ExecContext(ctx, query1)
	if err != nil {
		return fmt.Errorf("Error initializing schema: %w", err)
	}

	return nil
}
