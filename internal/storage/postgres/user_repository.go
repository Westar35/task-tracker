package postgres

import (
	"database/sql"
	"task-tracker/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, email, password_hash, created_at
	FROM users
	WHERE email = $1
	`

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, err
	}

	return user, nil
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
