package dto

import (
	"github.com/google/uuid"
	"time"
)

// TODO: 2. adjust DTO for request & response

type ExampleRequestDTO struct {
	Name        string  `json:"name" validate:"required"`
	ExampleType string  `json:"example_type" validate:"required,oneof=credit debit"`
	Amount      float64 `json:"amount" validate:"required,min=1"`
}

type ExampleResponseDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ExampleType string    `json:"example_type"`
	Amount      float64   `json:"amount"`
	Timestamp   time.Time `json:"timestamp"`
}

type GetExamplesResponseDTO struct {
	BaseResponse
	Data []ExampleResponseDTO `json:"data"`
}

type GetExampleResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
}

type CreateExampleResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
}

type UpdateExampleResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
}

type DeleteExampleResponseDTO struct {
	BaseResponse
}
