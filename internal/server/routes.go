package server

import (
	"net/http"
	"task-tracker/internal/handlers"
)

func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/tasks/", handlers.TaskByIDHandler)
}
