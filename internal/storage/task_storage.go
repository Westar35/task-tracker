package storage

import (
	"fmt"
	"task-tracker/internal/models"
)

// TaskStorage представляет срез задач и следующий ID для новой задачи.
type TaskStorage struct {
	tasks  []models.Task
	nextID int
}

// NewTaskStorage создает новый экземпляр TaskStorage с инициализированными полями.
func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

// GetAllTasks возвращает все задачи из хранилища.
func (s *TaskStorage) GetAllTasks() []models.Task {
	return s.tasks
}

// Не реализовано
func (s *TaskStorage) GetTaskByID(id int) (models.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, fmt.Errorf("task not found")
}

func (s *TaskStorage) CreateTask(title string) models.Task {
	return models.Task{
		ID:    s.nextID,
		Title: title,
		Done:  false,
	}
}

func (s *TaskStorage) AddTask(task models.Task) {
	s.tasks = append(s.tasks, task)
	s.nextID++
}

func (s *TaskStorage) DeleteTask(id int) error {
	return nil
}

func (s *TaskStorage) UpdateTask(id int, title string) error {
	return nil
}
