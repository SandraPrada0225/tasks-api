package services

import (
	"net/http"
	"tasks-api/internal/models"
	"tasks-api/internal/repository"
	"tasks-api/internal/utils"
)

//service valida si los datos tienen sentido en el sistema

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
	rowsAffected, err := s.repo.Update(id, title)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.NewError(
			"TASK_NOT_FOUND",
			"La tarea no existe",
			http.StatusNotFound,
		)
	}

	return nil
}

// Delete
func (s *TaskService) DeleteTask(id int) error {
	rowsAffected, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.NewError(
			"TASK_NOT_FOUND",
			"La tarea no existe",
			http.StatusNotFound,
		)
	}

	return nil
}
