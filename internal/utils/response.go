package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Data  interface{} `json:"data"`
	Error *AppError   `json:"error"`
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{
		Data:  data,
		Error: nil,
	}

	json.NewEncoder(w).Encode(resp)
}

func JSONError(w http.ResponseWriter, appErr *AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Status)

	resp := ErrorResponse{
		Data:  nil,
		Error: appErr,
	}

	json.NewEncoder(w).Encode(resp)
}
