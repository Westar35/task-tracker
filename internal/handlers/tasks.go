package handlers

import (
	"net/http"
	"task-tracker/internal/storage"
	"task-tracker/internal/models"
	"encoding/json"
)

var taskStorage = storage.NewTaskStorage()

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

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req models.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && req.Title != "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newTask := taskStorage.CreateTask(req.Title)
	json.NewEncoder(w).Encode(newTask)
	taskStorage.AddTask(newTask)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(taskStorage.GetAllTasks())
}