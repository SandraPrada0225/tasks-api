package middleware

import (
	"net/http"
	"tasks-api/internal/utils"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorMiddleware(next AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := next(w, r)
		if err != nil {
			//error controlado
			if appErr, ok := err.(*utils.AppError); ok {
				utils.JSONError(w, appErr.Status, appErr.Message)
				return
			}
			//error genérico
			utils.JSONError(w, http.StatusInternalServerError, "Error interno del servidor")
		}
		/*//validacion
		if valErr, ok := err.(*utils.ValidationError); ok {
			utils.JSONResponse(w, http.StatusBadRequest, valErr)
			return
		}
		//error personalizado
		if appError, ok := err.(*utils.AppError); ok {
			utils.JSONResponse(w, appError.Status, map[string]interface{}{
				"error": map[string]string{
					"code":    appError.Code,
					"message": appError.Message,
				},
			})
			return
		}*/

	}
}
