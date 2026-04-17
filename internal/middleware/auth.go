package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"task-tracker/internal/service"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "userID"

func AuthMiddleware(jwtSecret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			tokenString := parts[1]

			claims := &service.Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			userID, err := strconv.ParseInt(claims.Subject, 10, 64)
			if err != nil {
				http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(userIDKey).(int64)
	if !ok {
		return 0, fmt.Errorf("user ID not found in context")
	}
	return userID, nil
}
