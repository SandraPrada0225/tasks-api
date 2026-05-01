package middleware

import (
	"net/http"
	"tasks-api/internal/utils"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorMiddleware(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := h(w, r)
		if err == nil {
			return
		}
		//validacion
		if valErr, ok := err.(*utils.ValidationError); ok {
			utils.JSONResponse(w, http.StatusBadRequest, valErr)
			return
		}
		//error personalizado
		if appError, ok := err.(*utils.AppError); ok {
			utils.JSONError(w, appError.Code, appError.Message)
			return
		}
		//error genérico
		utils.JSONError(w, http.StatusInternalServerError, "Error interno del servidor")
		//fmt.Printf("Tipo de error: %T\n", err)
	}
}
