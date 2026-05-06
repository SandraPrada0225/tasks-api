package middleware

import (
	"log"
	"net/http"
	"tasks-api/internal/utils"
	"time"
)

type AppHandler func(http.ResponseWriter, *http.Request) error

func ErrorMiddleware(next AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		err := next(w, r)
		requestID := GetRequestID(r)

		duration := time.Since(start)
		if err != nil {
			//error controlado
			if appErr, ok := err.(*utils.AppError); ok {

				log.Printf(
					"[ERROR] [%s] %s %s |%d|%s|%v",
					requestID,
					r.Method,
					r.URL.Path,
					appErr.Status,
					appErr.Message,
					duration,
				)
				utils.JSONError(w, appErr)
				return
			}

			//error inesperado
			log.Printf(
				"[ERROR] %s %s |500|%v|%v",
				r.Method,
				r.URL.Path,
				err,
				duration,
			)

			//error genérico
			utils.JSONError(w, &utils.AppError{
				Code:    "INTERNAL_ERROR",
				Message: "Error interno",
				Status:  http.StatusInternalServerError,
			})
		}

		//exito
		log.Printf(
			"[OK] [%s] %s %s |%v",
			requestID,
			r.Method,
			r.URL.Path,
			duration,
		)

	}
}
