package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-tracker/internal/models"
	"task-tracker/internal/storage"
)

// TaskHandler - структура для хранения ссылки на TaskStorage и обработки HTTP-запросов, связанных с задачами.
type TaskHandler struct {
	taskPointer *storage.TaskStorage
}

// NewTaskHandler - конструктор для создания нового экземпляра TaskHandler с переданной ссылкой на TaskStorage.
func NewTaskHandler(taskStorage *storage.TaskStorage) *TaskHandler {
	return &TaskHandler{
		taskPointer: taskStorage,
	}
}

// TasksHandler - метод для обработки HTTP-запросов на эндпоинте /tasks.
func (h *TaskHandler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getAllTasks(w, r)
	case http.MethodPost:
		h.createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// TaskByIDHandler - метод для обработки HTTP-запросов на эндпоинте /tasks/{id}.
func (h *TaskHandler) TasksByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTaskByID(w, r)
	case http.MethodPut:
		//h.updateTask(w, r)
	case http.MethodDelete:
		//h.deleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
Методы для TasksHandler.
*/

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
	var req models.CreateTaskRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task := h.taskPointer.CreateTask(req.Title)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

/*
Методы для getTaskByID.
*/

func (h *TaskHandler) getTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	splittedPath := strings.Split(r.URL.Path, "/")
	if len(splittedPath) != 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	taskID, err := strconv.Atoi(splittedPath[2])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	task, err := h.taskPointer.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}
