package handlers

import (
	"net/http"
	"encoding/json"
	"task-tracker/internal/storage"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.NewHealthStorage())
	w.WriteHeader(http.StatusOK)
}