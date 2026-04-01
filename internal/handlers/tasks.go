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

// getIDFromPath - вспомогательная функция для извлечения ID задачи из URL-пути запроса.
func getIDFromPath(r *http.Request) int {
	splittedPath := strings.Split(r.URL.Path, "/")
	if len(splittedPath) != 3 {
		return 0
	}
	taskID, err := strconv.Atoi(splittedPath[2])
	if err != nil {
		return 0
	}
	return taskID
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
		h.deleteTaskById(w, r)
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
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	task := h.taskPointer.CreateTask(req.Title)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
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
	id := getIDFromPath(r)
	task, err := h.taskPointer.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) deleteTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := getIDFromPath(r)
	err := h.taskPointer.DeleteTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
