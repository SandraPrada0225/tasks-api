package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tasks-api/internal/dto"
	"tasks-api/internal/models"
	"tasks-api/internal/services"
	"tasks-api/internal/utils"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(s *services.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) error {

	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return utils.NewError("Error al obtener tareas", http.StatusInternalServerError)
	}
	utils.JSONResponse(w, http.StatusOK, tasks)
	return nil
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) error {
	var body dto.CreateTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return utils.NewError("Error al leer datos", http.StatusBadRequest)
	}
	//validaciones

	if err := body.Validate(); err != nil {
		return err
	}

	task, err := h.service.CreateTask(body.Title)
	if err != nil {
		return utils.NewError("Error al guadar los datos en la DB", http.StatusInternalServerError)
	}
	//responde con la tarea creada
	utils.JSONResponse(w, http.StatusCreated, task)
	return nil
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.NewError("ID invalido", http.StatusBadRequest)
	}

	var UpdateTask models.Task
	err = json.NewDecoder(r.Body).Decode(&UpdateTask)
	if err != nil {
		return utils.NewError(err.Error(), http.StatusBadRequest)
	}

	//validaciones
	if strings.TrimSpace(UpdateTask.Title) == "" {
		return utils.NewError("El titulo es obligatorio", http.StatusBadRequest)
	}

	err = h.service.UpdateTask(id, UpdateTask.Title)
	if err != nil {
		return utils.NewError(err.Error(), http.StatusNotFound)
	}
	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Tarea actualizada",
	})
	return nil
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.NewError("ID invalido", http.StatusBadRequest)
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		return utils.NewError(err.Error(), http.StatusNotFound)
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Tarea eliminada",
	})
	return nil
}
