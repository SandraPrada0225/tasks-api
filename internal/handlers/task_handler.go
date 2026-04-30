package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	//ejecutamos el query
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error al obtener los datos")
		return
	}
	utils.JSONResponse(w, http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "Error al leer datos")
		return
	}
	//validaciones
	if strings.TrimSpace(newTask.Title) == "" {
		utils.JSONError(w, http.StatusBadRequest, "El titulo es obligatorio")
		return
	}

	if len(newTask.Title) > 100 {
		utils.JSONError(w, http.StatusBadRequest, "El titulo es demasiado largo")
		return
	}

	task, err := h.service.CreateTask(newTask.Title)
	if err != nil {
		utils.JSONError(w, http.StatusInternalServerError, "Error al guadar los datos en la DB")
		return
	}
	//responde con la tarea creada
	utils.JSONResponse(w, http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	var UpdateTask models.Task
	err = json.NewDecoder(r.Body).Decode(&UpdateTask)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	//validaciones
	if strings.TrimSpace(UpdateTask.Title) == "" {
		utils.JSONError(w, http.StatusBadRequest, "El titulo es obligatorio")
		return
	}

	err = h.service.UpdateTask(id, UpdateTask.Title)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Tarea actualizada",
	})
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.JSONError(w, http.StatusBadRequest, "ID invalido")
		return
	}

	err = h.service.DeleteTask(id)
	if err != nil {
		utils.JSONError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Tarea eliminada",
	})

}
