package middleware

import (
	"log"
	"net/http"
)


func RecoveryMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		
		defer func() {
			if err := recover(); err != nil {
				log.Println("PANIC:", err)

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"Error interno" }`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}