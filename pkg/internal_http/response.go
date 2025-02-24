package internal_http

import (
	"encoding/json"
	"net/http"

	"github.com/irfan44/go-http-boilerplate/internal/dto"
)

func NewBaseResponse(message string) dto.BaseResponse {
	return dto.BaseResponse{
		Message: message,
	}
}

func ApplyJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SendResponse(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
