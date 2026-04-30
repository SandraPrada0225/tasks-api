package services

import (
	"tasks-api/internal/models"
	"tasks-api/internal/repository"
)

// estructura
type TaskService struct {
	repo repository.TaskRepository
}

// constructor
func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

// ver
func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

// crear
func (s *TaskService) CreateTask(title string) (models.Task, error) {
	return s.repo.Create(title)
}

// update
func (s *TaskService) UpdateTask(id int, title string) error {
	return s.repo.Update(id, title)
}

// Delete
func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
