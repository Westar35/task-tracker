package postgres

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(email, password string) error {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(query, email, password)
	if err != nil {
		return err
	}

	return nil
}
