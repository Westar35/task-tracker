package server

import (
	"net/http"
	"task-tracker/internal/handlers"
)

func registerRoutes(mux *http.ServeMux, taskHandler *handlers.TaskHandler) {
	// mux.HandleFunc("/", handlers.RootHandler)
	// mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/tasks", taskHandler.Tasks)
	//mux.HandleFunc("/tasks/", handlers.TaskByIDHandler)
}
