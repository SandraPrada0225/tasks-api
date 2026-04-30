package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) writer(b []byte) (int, error) {
	if rw.statusCode == 0 {
		rw.statusCode = 200
	}
	return rw.ResponseWriter.Write(b)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		//envolvemos el writer
		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     0, //default
		}

		next.ServeHTTP(rw, r)

		fmt.Println(
			r.Method,
			r.URL.Path,
			"→",
			rw.statusCode,
			"→",
			time.Since(start))
	})
}
