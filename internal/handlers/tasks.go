package handlers

import (
	"encoding/json"
	"fmt"
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

func getIDFromURL(path string) (int, error) {
	var id int
	_, err := fmt.Sscanf(path, "/tasks/%d", &id)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID in URL: %v", err)
	}
	return id, nil
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
		id, err := getIDFromURL(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		h.GetTaskByID(id, w, r)
	case http.MethodPut:
		id, err := getIDFromURL(r.URL.Path)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		h.UpdateTaskByID(id, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
Методы для TasksHandler
*/

// Метод GetAllTasks получает все задачи из репозитория и возвращает их в виде JSON-ответа
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

// Метод CreateTask создает новую задачу на основе данных из запроса и сохраняет ее в базе данных
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
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}
}

/*
Методы для TasksByIDHandler
*/

func (h *TaskHandler) GetTaskByID(id int, w http.ResponseWriter, r *http.Request) {
	task, err := h.repo.GetTaskByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Task with current ID not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTaskByID(id int, w http.ResponseWriter, r *http.Request) {
	var task models.UpdateTaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.repo.UpdateTaskByID(id, task.Title, task.Status)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(updatedTask)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to encode updated task", http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) DeleteTaskByID(id int, w http.ResponseWriter, r *http.Request) {}
