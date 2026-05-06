package middleware

import "net/http"

func GetRequestID(r *http.Request) string {
	if id, ok := r.Context().Value(requestIDKey).(string); ok {
		return id
	}
	return "unknown"
}
