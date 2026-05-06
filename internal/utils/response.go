package utils

import (
	"encoding/json"
	"net/http"
)

type APIresponse struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIresponse{
		Data:  data,
		Error: nil,
	}

	json.NewEncoder(w).Encode(resp)
}

func JSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := APIresponse{
		Data:  nil,
		Error: message,
	}

	json.NewEncoder(w).Encode(resp)
}
