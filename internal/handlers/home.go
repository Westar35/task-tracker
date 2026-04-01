package handlers

import (
	"encoding/json"
	"net/http"
	"task-tracker/internal/storage"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(storage.NewHome())
}
