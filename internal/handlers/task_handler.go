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
	service services.TaskServiceInterface
}

func NewTaskHandler(s services.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) error {

	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return err
	}
	utils.JSONResponse(w, http.StatusOK, tasks)
	return nil
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return utils.ErrInvalidID()
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return err
	}

	utils.JSONResponse(w, http.StatusOK, task)
	return nil
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) error {
	var body dto.CreateTaskDTO

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return utils.NewAppError(
			"INVALID_JSON",
			"JSON inválido",
			http.StatusBadRequest)
	}
	//validaciones

	if err := body.Validate(); err != nil {
		// detectar ValidationError
		if ve, ok := err.(*utils.ValidationError); ok {
			return utils.NewAppError(
				"VALIDATION_ERROR",
				ve.Error(), // usa el mensaje construido
				http.StatusBadRequest,
			)
		}
		return err
	}

	task, err := h.service.CreateTask(body.Title)
	if err != nil {
		return err
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
		return utils.ErrInvalidID()
	}

	var UpdateTask models.Task
	err = json.NewDecoder(r.Body).Decode(&UpdateTask)
	if err != nil {
		return utils.NewAppError(
			"INAVLID_JSON",
			"JSON invalido",
			http.StatusBadRequest)
	}

	//validaciones
	if strings.TrimSpace(UpdateTask.Title) == "" {
		return utils.NewAppError(
			"INAVLID_TITLE",
			"El titulo es obligatorio",
			http.StatusBadRequest)
	}

	err = h.service.UpdateTask(id, UpdateTask.Title)
	if err != nil {
		return err
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
		return utils.ErrInvalidID()
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		return err
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Tarea eliminada",
	})
	return nil
}
