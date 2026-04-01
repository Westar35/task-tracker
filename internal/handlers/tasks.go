package handlers

import (
	"encoding/json"
	"net/http"
	"task-tracker/internal/storage"
)

type TaskHandler struct {
	taskPointer *storage.TaskStorage
}

func NewTaskHandler(taskStorage *storage.TaskStorage) *TaskHandler {
	return &TaskHandler{
		taskPointer: taskStorage,
	}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllTasks(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(h.taskPointer.GetAllTasks())
}

func (h *TaskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	h.taskPointer.CreateTask()
	json.NewEncoder(w).Encode(h)
}
