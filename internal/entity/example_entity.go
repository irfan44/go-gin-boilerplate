package entity

import (
	"github.com/google/uuid"
	"github.com/irfan44/go-http-boilerplate/internal/dto"
	"time"
)

// TODO: 1. adjust entity based on model

type Example struct {
	Id          uuid.UUID
	Name        string
	ExampleType string
	Amount      float64
	Timestamp   time.Time
}

func (e *Example) ToExampleDTO() *dto.ExampleResponseDTO {
	return &dto.ExampleResponseDTO{
		Id:          e.Id,
		Name:        e.Name,
		ExampleType: e.ExampleType,
		Amount:      e.Amount,
		Timestamp:   e.Timestamp,
	}
}

type Examples []Example

func (e Examples) ToExamplesDTO() []dto.ExampleResponseDTO {
	var result []dto.ExampleResponseDTO

	for _, example := range e {
		result = append(result, *example.ToExampleDTO())
	}

	return result
}
