package internal_http

import (
	"github.com/irfan44/go-gin-boilerplate/internal/dto"
)

func NewBaseResponse(message string) dto.BaseResponse {
	return dto.BaseResponse{
		Message: message,
	}
}
