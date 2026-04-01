package handlers

import (
	"encoding/json"
	"net/http"
	"task-tracker/internal/storage"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(storage.NewHome())
	if err != nil {
		http.Error(w, "Failed to encode home status", http.StatusInternalServerError)
	}
}
