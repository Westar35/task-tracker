package server

import (
	"fmt"
	"log"
	"net/http"
	"task-tracker/internal/handlers"
	"task-tracker/internal/storage"
)

func Run() error {
	mux := http.NewServeMux()

	taskStorage := storage.NewTaskStorage()
	taskHandler := handlers.NewTaskHandler(taskStorage)

	registerRoutes(mux, taskHandler)

	log.Println("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("start server %w", err)
	}
	return nil
}
