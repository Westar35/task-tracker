package server

import (
	"net/http"
	"task-tracker/internal/handlers"
)

func registerRoutes(mux *http.ServeMux, taskHandler *handlers.TaskHandler, authHandler *handlers.AuthHandler) {
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/tasks", taskHandler.TasksHandler)
	mux.HandleFunc("/tasks/", taskHandler.TasksByIDHandler)
	mux.HandleFunc("/auth/register", authHandler.RegisterHandler)
	mux.HandleFunc("/auth/login", authHandler.LoginHandler)
}
