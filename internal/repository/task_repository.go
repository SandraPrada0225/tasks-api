package repository

import "tasks-api/internal/models"

// definimos un contrato
type TaskRepository interface {
	GetAll() ([]models.Task, error)
	Create(title string) (models.Task, error)
	Update(id int, title string) (int64, error)
	Delete(id int) (int64, error)
}
