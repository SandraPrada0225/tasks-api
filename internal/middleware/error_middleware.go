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

		duration := time.Since(start)
		if err != nil {
			//error controlado
			if appErr, ok := err.(*utils.AppError); ok {

				log.Printf(
					"[ERROR] %s %s |%d|%s|%v",
					r.Method,
					r.URL.Path,
					appErr.Status,
					appErr.Message,
					duration,
				)
				utils.JSONError(w, appErr.Status, appErr.Message)
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
			utils.JSONError(w, http.StatusInternalServerError, "Error interno del servidor")
		}

		//exito
		log.Printf(
			"[OK] %s %s |%v",
			r.Method,
			r.URL.Path,
			duration,
		)

	}
}
