package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"tasks-api/internal/middleware"
	"tasks-api/internal/models"
	"tasks-api/internal/utils"
	"testing"

	"github.com/gorilla/mux"
)

//no usamos service, lo mockeamos

type mockService struct {
	getByIDFunc func(in int) (models.Task, error)
	createFunc  func(title string) (models.Task, error)
	updateFunc  func(id int, title string) error
	deleteFunc  func(id int) error
	getAllFunc  func() ([]models.Task, error)
}
type ErrorResponse struct {
	Error string `json:"error"`
}
type Response struct {
	Data models.Task `json:"data"`
}

func (m *mockService) GetTaskByID(id int) (models.Task, error) {
	return m.getByIDFunc(id)
}
func (m *mockService) GetAllTasks() ([]models.Task, error) {
	return m.getAllFunc()
}

func (m *mockService) CreateTask(title string) (models.Task, error) {
	return m.createFunc(title)
}

func (m *mockService) UpdateTask(id int, title string) error {
	return m.updateFunc(id, title)
}

func (m *mockService) DeleteTask(id int) error {
	return m.deleteFunc(id)
}

func TestCreateTask_Handler(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockFunc       func(title string) (models.Task, error)
		expectedStatus int
	}{
		{
			name: "succes",
			body: `{"title": "Nueva tarea"}`,
			mockFunc: func(title string) (models.Task, error) {
				return models.Task{ID: 1, Title: title}, nil
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid json",
			body:           `{"title":`,
			mockFunc:       nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "empty title",
			body: `{"title":""}`,
			mockFunc: func(title string) (models.Task, error) {
				return models.Task{}, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			service := &mockService{}

			if tt.mockFunc != nil {
				service.createFunc = tt.mockFunc
			}
			handler := &TaskHandler{
				service: service,
			}
			req := httptest.NewRequest("POST", "/tasks", strings.NewReader(tt.body))
			req.Header.Set("Content-Typre", "application/json")

			rr := httptest.NewRecorder()

			wrapped := middleware.ErrorMiddleware(handler.CreateTask)
			wrapped(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.name == "succes" {
				var resp Response

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Data.ID != 1 {
					t.Errorf("expected ID 1, got %d", resp.Data.ID)
				}

				if resp.Data.Title != "Nueva tarea" {
					t.Errorf("unexpected title: %s", resp.Data.Title)
				}
			}
			if tt.name == "invalid json" || tt.name == "empty title" {

				var resp ErrorResponse

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Error == "" {
					t.Errorf("expected error message")
				}
			}

		})
	}
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

			wrapped := middleware.ErrorMiddleware(handler.GetTaskByID)
			wrapped(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.name == "succes" {
				var resp Response

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Data.ID != 1 {
					t.Errorf("expected ID 1, got %d", resp.Data.ID)
				}

				if resp.Data.Title != "Test" {
					t.Errorf("expected title Test, got %s", resp.Data.Title)
				}
			}
			if tt.name == "not found" || tt.name == "invalid id" {

				var resp ErrorResponse

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Error == "" {
					t.Errorf("expected error message, got empty")
				}
			}

		})
	}

}

func TestUpdate_Handler(t *testing.T) {

	tests := []struct {
		name           string
		id             string
		body           string
		mockFunc       func(id int, title string) error
		expectedStatus int
	}{
		{
			name: "succes",
			id:   "1",
			body: `{"title":"Actualizada"}`,
			mockFunc: func(id int, title string) error {
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			id:             "abc",
			body:           `{"title":"Test"}`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "empty title",
			id:   "1",
			body: `{"title":""}`,
			mockFunc: func(id int, title string) error {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "not found",
			id:   "999",
			body: `{"title":"Test"}`,
			mockFunc: func(id int, title string) error {
				return utils.ErrTaskNotFound()
			},
			expectedStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			service := &mockService{}

			if tt.mockFunc != nil {
				service.updateFunc = tt.mockFunc
			}
			handler := &TaskHandler{
				service: service,
			}
			req := httptest.NewRequest("PUT", "/tasks/"+tt.id, strings.NewReader(tt.body))
			vars := map[string]string{
				"id": tt.id,
			}
			req = mux.SetURLVars(req, vars)

			rr := httptest.NewRecorder()

			wrapped := middleware.ErrorMiddleware(handler.UpdateTask)
			wrapped(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.name == "succes" {
				var resp Response

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Data.Title != "" {
					t.Errorf("expected message, got empty")
				}
			}
			if tt.name == "not found" || tt.name == "invalid id" {

				var resp ErrorResponse

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Error == "" {
					t.Errorf("expected error message, got empty")
				}
			}

		})
	}

}

func TestDelete_Handler(t *testing.T) {

	tests := []struct {
		name           string
		id             string
		mockFunc       func(id int) error
		expectedStatus int
	}{
		{
			name: "succes",
			id:   "1",
			mockFunc: func(id int) error {
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid id",
			id:             "abc",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "not found",
			id:   "999",
			mockFunc: func(id int) error {
				return utils.ErrTaskNotFound()
			},
			expectedStatus: http.StatusNotFound,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			service := &mockService{}

			if tt.mockFunc != nil {
				service.deleteFunc = tt.mockFunc
			}
			handler := &TaskHandler{
				service: service,
			}
			req := httptest.NewRequest("DELETE", "/tasks/"+tt.id, nil)

			vars := map[string]string{
				"id": tt.id,
			}
			req = mux.SetURLVars(req, vars)

			rr := httptest.NewRecorder()

			wrapped := middleware.ErrorMiddleware(handler.DeleteTask)
			wrapped(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.name == "succes" {
				var resp Response

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Data.Title != "" {
					t.Errorf("expected message, got empty")
				}
			}
			if tt.name == "not found" || tt.name == "invalid id" {

				var resp ErrorResponse

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Error == "" {
					t.Errorf("expected error message, got empty")
				}
			}
		})
	}
}

func TestGetAllTask_Handler(t *testing.T) {

	tests := []struct {
		name           string
		mockFunc       func() ([]models.Task, error)
		expectedStatus int
		expectedLen    int
	}{
		{
			name: "success",
			mockFunc: func() ([]models.Task, error) {
				return []models.Task{
					{ID: 1, Title: "Task 1"},
					{ID: 2, Title: "Task 2"},
				}, nil
			},
			expectedStatus: http.StatusOK,
			expectedLen:    2,
		},
		{
			name: "empty list",
			mockFunc: func() ([]models.Task, error) {
				return []models.Task{}, nil
			},
			expectedStatus: http.StatusOK,
			expectedLen:    0,
		},
		{
			name: "service error",
			mockFunc: func() ([]models.Task, error) {
				return nil, utils.NewAppError(
					"INTERNAL_ERROR",
					"error al obtener tareas",
					http.StatusInternalServerError,
				)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			service := &mockService{
				getAllFunc: tt.mockFunc,
			}

			handler := &TaskHandler{
				service: service,
			}
			req := httptest.NewRequest("GET", "/tasks", nil)
			rr := httptest.NewRecorder()

			wrapped := middleware.ErrorMiddleware(handler.GetTasks)
			wrapped(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.name == "succes" || tt.name == "empty list" {

				type Response struct {
					Data []models.Task `json:"data"`
				}
				var resp Response

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if len(resp.Data) != tt.expectedLen {
					t.Errorf("expected %d tasks, got %d", tt.expectedLen, len(resp.Data))
				}
			}
			if tt.name == "service error" {

				var resp ErrorResponse

				err := json.Unmarshal(rr.Body.Bytes(), &resp)
				if err != nil {
					t.Errorf("error parsing JSON: %v", err)
				}

				if resp.Error == "" {
					t.Errorf("expected error message, got empty")
				}
			}
		})
	}
}
