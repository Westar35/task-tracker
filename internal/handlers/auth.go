package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"task-tracker/internal/models"
	"task-tracker/internal/service"
	"task-tracker/internal/storage/postgres"
)

type AuthHandler struct {
	repo *postgres.UserRepository
}

func NewAuthHandler(repo *postgres.UserRepository) *AuthHandler {
	return &AuthHandler{repo: repo}
}

var jwtSecret = []byte("dev-task-tracker-secret-key-2026")

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.RegisterUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.LoginUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("RegisterHandler error: %v", err)
	}

	err = service.CheckUserRegisterInput(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Email = service.EmailToLower(req.Email)
	err = service.CheckInputPassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req.Password, err = service.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		log.Printf("RegisterHandler error: %v", err)
		return
	}

	err = h.repo.CreateUser(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("RegisterHandler error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("LoginHandler error: %v", err)
		return
	}

	req.Email = service.EmailToLower(req.Email)
	err = service.CheckUserRegisterInput(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByEmail(req.Email)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		log.Printf("LoginHandler error: %v", err)
		return
	}

	err = service.ComparePassword(user.PasswordHash, req.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := service.GenerateAccessToken(user.ID, jwtSecret)
	log.Println("Generated JWT:", token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}
