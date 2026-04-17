package service

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

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
