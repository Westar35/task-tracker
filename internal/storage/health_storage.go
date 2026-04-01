package storage

import (
	"task-tracker/internal/models"
)

func NewHealthStorage() *models.Health {
	return &models.Health{
		Status: "ok",
	}
}