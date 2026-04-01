package handlers

import (
	"net/http"
	"encoding/json"
	"task-tracker/internal/storage"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.NewHome())
	w.WriteHeader(http.StatusOK)
}