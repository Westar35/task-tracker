package storage

import (
	"task-tracker/internal/models"
)

type TaskStorage struct {
	tasks  []models.Task
	nextID int
}

func NewTaskStorage() *TaskStorage {
	return &TaskStorage{
		tasks:  []models.Task{},
		nextID: 1,
	}
}

func (s *TaskStorage) GetAllTasks() []models.Task {
	return s.tasks
}

// Не реализовано
func (s *TaskStorage) GetTaskByID(id int) (*models.Task, error) {
	return &models.Task{}, nil
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
}

func (s *TaskStorage) DeleteTask(id int) error {
	return nil
}

func (s *TaskStorage) UpdateTask(id int, title string) error {
	return nil
}
