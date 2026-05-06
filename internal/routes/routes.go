package routes

import (
	"tasks-api/internal/handlers"
	"tasks-api/internal/middleware"

	"github.com/gorilla/mux"
)

// funcion que registra los endpoints
func RegisterRoutes(handler *handlers.TaskHandler) *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.RequestIDMiddleware)

	r.HandleFunc("/tasks", middleware.ErrorMiddleware(handler.GetTasks)).Methods("GET")
	r.HandleFunc("/tasks/{id}", middleware.ErrorMiddleware(handler.GetTaskByID)).Methods("GET")
	r.HandleFunc("/tasks", middleware.ErrorMiddleware(handler.CreateTask)).Methods("POST")
	r.HandleFunc("/tasks/{id}", middleware.ErrorMiddleware(handler.UpdateTask)).Methods("PUT")
	r.HandleFunc("/tasks/{id}", middleware.ErrorMiddleware(handler.DeleteTask)).Methods("DELETE")

	return r
}
