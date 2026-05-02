package utils

//Si algo se repite y no pertenece a una capa específica → utils
import "net/http"

func ErrInvalidID() *AppError {
	return NewAppError("INVALID_ID", "ID inválido", http.StatusBadRequest)
}

func ErrTaskNotFound() *AppError {
	return NewAppError("TASK_NOT_FOUND", "La tarea no existe", http.StatusNotFound)
}

func ErrInternal() *AppError {
	return NewAppError("INTERNAL_ERROR", "Error interno", http.StatusInternalServerError)
}
