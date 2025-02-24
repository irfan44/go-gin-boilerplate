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
} // @name ExampleRequest

type ExampleResponseDTO struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	ExampleType string    `json:"example_type"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
} // @name ExampleResponse

type GetExamplesResponseDTO struct {
	BaseResponse
	Data []ExampleResponseDTO `json:"data"`
} // @name GetExamplesResponse

type GetExampleByIdResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
} // @name GetExampleByIdResponse

type CreateExampleResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
} // @name CreateExampleResponse

type UpdateExampleResponseDTO struct {
	BaseResponse
	Data ExampleResponseDTO `json:"data"`
} // @name UpdateExampleResponse

type DeleteExampleResponseDTO struct {
	BaseResponse
} // @name DeleteExampleResponse
