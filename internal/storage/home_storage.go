package storage

import (
	"task-tracker/internal/models"
)

func NewHome() *models.Home {
	return &models.Home{
		Title: "Server is running",
	}
}