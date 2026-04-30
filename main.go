package main

import (
	"fmt"
	"net/http"
	"tasks-api/internal/database"
	"tasks-api/internal/handlers"
	"tasks-api/internal/repository"
	"tasks-api/internal/routes"
	"tasks-api/internal/services"
)

func main() {
	//abre la conexion y verifica con ping
	database.Connect()
	repo := &repository.PostgresTaskRepository{
		DB: database.DB,
	}
	service := services.NewTaskService(repo)

	handler := handlers.NewTaskHandler(service)
	//registra todo los endpoints
	router := routes.RegisterRoutes(handler)
	//arranca el serivor
	fmt.Println("Servidor en http://localhost:8080")
	http.ListenAndServe(":8080", router)
}

/* main.go → arranca servidor
routes → define endpoints
handlers → lógica
models → datos*/
