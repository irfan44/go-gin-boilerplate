package internal_http

import (
	"github.com/irfan44/go-http-boilerplate/internal/dto"
)

func NewBaseResponse(message string) dto.BaseResponse {
	return dto.BaseResponse{
		Message: message,
	}
}
