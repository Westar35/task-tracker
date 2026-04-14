package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"task-tracker/internal/handlers"
	"task-tracker/internal/storage/postgres"
)

func Run(db *sql.DB) error {
	mux := http.NewServeMux()

	taskRepo := postgres.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	registerRoutes(mux, taskHandler)

	log.Println("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("start server %w", err)
	}
	return nil
}
