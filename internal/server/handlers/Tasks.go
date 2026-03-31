package handlers

import (
	//"encoding/json"
	"net/http"
	//"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasksHandler(w, r)
	case http.MethodPost:
		createTaskHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
