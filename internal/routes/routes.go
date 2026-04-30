package routes

import (
	"tasks-api/internal/handlers"
	"tasks-api/internal/middleware"

	"github.com/gorilla/mux"
)

// funcion que registra los endpoints
func RegisterRoutes(handler *handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.Logger)

	r.HandleFunc("/tasks", handler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")

	return r
}
