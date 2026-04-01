package handlers

import (
	"encoding/json"
	"net/http"
	"task-tracker/internal/storage"
)

type TaskHandler struct {
	storage *storage.TaskStorage
}

func NewTaskHandler(taskStorage *storage.TaskStorage) *TaskHandler {
	return &TaskHandler{
		storage: taskStorage,
	}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllTasks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(h)
}

//123456
