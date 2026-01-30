package helpers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	status := "success"
	if code >= 400 {
		status = "error"
	}

	response := APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func JSONError(w http.ResponseWriter, code int, message string) {
	JSONResponse(w, code, message, nil)
}
