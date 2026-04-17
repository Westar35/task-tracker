package service

import (
	"fmt"
	"strconv"
	"strings"
	"task-tracker/internal/storage/postgres"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo  *postgres.UserRepository
	jwtSecret []byte
}

func NewAuthService(userRepo *postgres.UserRepository, jwtSecret []byte) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) GenerateJWT(userID int64) (string, error) {
	return "", nil
}

func CheckUserRegisterInput(email, password string) error {
	if email == "" {
		return fmt.Errorf("email is required")
	}
	if password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func EmailToLower(email string) string {
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)
	return email
}

func CheckInputPassword(password string) error {
	passwordRunes := []rune(password)
	if len(passwordRunes) < 8 || len(passwordRunes) > 18 {
		return fmt.Errorf("password must be between 8 and 18 characters long")
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hash), nil
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
}

func GenerateAccessToken(userID int64, secret []byte) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
