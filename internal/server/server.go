package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() error{
	mux := http.NewServeMux()
	registerRoutes(mux)

	log.Println("Starting server on port :8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return fmt.Errorf("start server %w", err)
	}
	return nil
}