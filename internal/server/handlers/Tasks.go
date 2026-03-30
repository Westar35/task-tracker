package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks := []Task{
			{ID: 1, Title: "Task 1"},
		}
		json.NewEncoder(w).Encode(tasks)
		//GetTasksHandler(w, r)
	case http.MethodPost:
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		fmt.Fprintf(w, "Создана задача: %s", task.Title)
		//CreateTaskHandler(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
