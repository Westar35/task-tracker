package server

import (
	"net/http"
	"task-tracker/internal/handlers"
)

func registerRoutes(mux *http.ServeMux, taskHandler *handlers.TaskHandler) {
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/tasks", taskHandler.TasksHandler)
	mux.HandleFunc("/tasks/", taskHandler.TasksByIDHandler)
}
