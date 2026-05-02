package services

//este es un repositorio falso, no usamos DB real, y controlamos lo que devuelve
import (
	"database/sql"
	"errors"
	"tasks-api/internal/models"
	"tasks-api/internal/utils"
	"testing"
)

type mockTaskRepo struct {
	getByIDFunc func(id int) (models.Task, error)
}

func (m *mockTaskRepo) GetByID(id int) (models.Task, error) {
	return m.getByIDFunc(id)
}

// implementamos lo minimo para cumplir la interfaz
func (m *mockTaskRepo) GetAll() ([]models.Task, error)             { return nil, nil }
func (m *mockTaskRepo) Create(title string) (models.Task, error)   { return models.Task{}, nil }
func (m *mockTaskRepo) Update(id int, title string) (int64, error) { return 0, nil }
func (m *mockTaskRepo) Delete(id int) (int64, error)               { return 0, nil }

func TestGetTaskByID(t *testing.T) {
	tests := []struct {
		name           string
		mockFunc       func(id int) (models.Task, error)
		inputID        int
		expectedError  string
		expectedTaskID int
	}{
		{
			name: "success",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{ID: 1, Title: "Test"}, nil
			},
			inputID:        1,
			expectedError:  "",
			expectedTaskID: 1,
		},
		{
			name: "not found",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{}, sql.ErrNoRows
			},
			inputID:        999,
			expectedError:  "TASK_NOT_FOUND",
			expectedTaskID: 0,
		},
		{
			name: "db error",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{}, errors.New("db error")
			},
			inputID:        1,
			expectedError:  "internal",
			expectedTaskID: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo := &mockTaskRepo{
				getByIDFunc: tt.mockFunc,
			}

			service := NewTaskService(mockRepo)

			task, err := service.GetTaskByID(tt.inputID)

			// validamos el error
			if tt.expectedError != "" {

				if err == nil {
					t.Errorf("expected error, go nil")
					return
				}

				//caso de negocio
				if tt.expectedError == "TASK_NOT_FOUND" {
					appErr, ok := err.(*utils.AppError)
					if !ok {
						t.Errorf("expected AppError, got %T", err)
						return
					}

					if appErr.Code != "TASK_NOT_FOUND" {
						t.Errorf("expected TASK_NOT_FOUND, got %s", appErr.Code)
					}
				}

				// caso error técnico
				if tt.expectedError == "internal" {
					if err == nil {
						t.Errorf("expected error, got nil")
					}
				}

				return
			}

			// validar éxito
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if task.ID != tt.expectedTaskID {
				t.Errorf("expected ID %d, got %d", tt.expectedTaskID, task.ID)
			}
		})
	}
}
