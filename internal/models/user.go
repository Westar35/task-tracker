package models

type User struct {
	ID           int64  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	EMail string `json:"email"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}
