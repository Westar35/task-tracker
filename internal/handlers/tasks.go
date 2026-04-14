package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"task-tracker/internal/models"
	"task-tracker/internal/storage/postgres"
)

type TaskHandler struct {
	repo *postgres.TaskRepository
}

func NewTaskHandler(repo *postgres.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// TasksHandler обрабатывает запросы к эндпоинту /tasks
func (h *TaskHandler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)
	case http.MethodGet:
		h.GetAllTasks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// TasksByIDHandler обрабатывает запросы к эндпоинту /tasks/{id}
func (h *TaskHandler) TasksByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Реализуйте логику получения задачи по ID
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
Методы для TasksHandler
*/

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.GetAllTasks()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tasksID := make([]int, 0, len(tasks))
	for id := range tasks {
		tasksID = append(tasksID, id)
	}
	slices.Sort(tasksID)
	for _, id := range tasksID {
		json.NewEncoder(w).Encode(tasks[id])
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	// Создаем экземпляр структуры запроса
	var req models.CreateTaskRequest

	// Декодируем JSON из тела запроса в структуру
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Проверяем, что заголовок не пустой
	if req.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Вызываем метод, который сохранит задачу в базе данных
	task, err := h.repo.CreateTask(req.Title)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

/*
Методы для TasksByIDHandler
*/
