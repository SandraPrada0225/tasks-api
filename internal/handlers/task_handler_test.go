package handlers

import (
	"net/http"
	"net/http/httptest"
	"tasks-api/internal/models"
	"tasks-api/internal/utils"
	"testing"

	"github.com/gorilla/mux"
)

//no usamos service, lo mockeamos

type mockService struct {
	getByIDFunc func(in int) (models.Task, error)
}

func (m *mockService) GetTaskByID(id int) (models.Task, error) {
	return m.getByIDFunc(id)
}
func (m *mockService) GetAllTasks() ([]models.Task, error) {
	return nil, nil
}

func (m *mockService) CreateTask(title string) (models.Task, error) {
	return models.Task{}, nil
}

func (m *mockService) UpdateTask(id int, title string) error {
	return nil
}

func (m *mockService) DeleteTask(id int) error {
	return nil
}

func TestGetByID_Handler(t *testing.T) {

	tests := []struct {
		name           string
		id             string
		mockFunc       func(id int) (models.Task, error)
		expectedStatus int
	}{
		{
			name: "succes",
			id:   "1",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{ID: 1, Title: "Test"}, nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "not found",
			id:   "999",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{}, utils.ErrTaskNotFound()
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "invalid id",
			id:   "abc",
			mockFunc: func(id int) (models.Task, error) {
				return models.Task{}, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			service := &mockService{
				getByIDFunc: tt.mockFunc,
			}
			handler := &TaskHandler{
				service: service,
			}
			req := httptest.NewRequest("GET", "/tasks/"+tt.id, nil)

			vars := map[string]string{
				"id": tt.id,
			}
			req = mux.SetURLVars(req, vars)

			rr := httptest.NewRecorder()

			err := handler.GetTaskByID(rr, req)

			if err != nil {
				appErr, ok := err.(*utils.AppError)
				if ok {
					rr.WriteHeader(appErr.Status)
				}
			}

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

		})
	}

}
