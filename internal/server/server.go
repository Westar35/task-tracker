package server

import (
	"fmt"
	"log"
	"net/http"
	"task-tracker/internal/server/handlers"
)

func Run() error{
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/tasks", handlers.TasksHandler)

	log.Println("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("start server %w", err)
	}
	return nil
}