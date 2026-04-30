package repository

import "tasks-api/internal/models"

// definimos un contrato
type TaskRepository interface {
	GetAll() ([]models.Task, error)
	Create(title string) (models.Task, error)
	Update(id int, title string) error
	Delete(id int) error
}
