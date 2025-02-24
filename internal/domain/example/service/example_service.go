package example_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/irfan44/go-http-boilerplate/internal/repository/example"
	"github.com/irfan44/go-http-boilerplate/pkg/errors"

	"github.com/irfan44/go-http-boilerplate/internal/dto"
)

type (
	ExampleService interface {
		GetExamples(ctx context.Context) (*dto.GetExamplesResponseDTO, errors.MessageErr)
		GetExampleById(id uuid.UUID, ctx context.Context) (*dto.GetExampleResponseDTO, errors.MessageErr)
		CreateExample(example dto.ExampleRequestDTO, ctx context.Context) (*dto.CreateExampleResponseDTO, errors.MessageErr)
		UpdateExample(example dto.ExampleRequestDTO, id uuid.UUID, ctx context.Context) (*dto.UpdateExampleResponseDTO, errors.MessageErr)
		DeleteExample(id uuid.UUID, ctx context.Context) (*dto.DeleteExampleResponseDTO, errors.MessageErr)
	}

	exampleService struct {
		repo example_repository.ExampleRepository
	}
)

// TODO: 4. adjust service

func (s *exampleService) GetExamples(ctx context.Context) (*dto.GetExamplesResponseDTO, errors.MessageErr) {
	return nil, nil
}

func (s *exampleService) GetExampleById(id uuid.UUID, ctx context.Context) (*dto.GetExampleResponseDTO, errors.MessageErr) {
	return nil, nil
}

func (s *exampleService) CreateExample(example dto.ExampleRequestDTO, ctx context.Context) (*dto.CreateExampleResponseDTO, errors.MessageErr) {
	return nil, nil
}

func (s *exampleService) UpdateExample(example dto.ExampleRequestDTO, id uuid.UUID, ctx context.Context) (*dto.UpdateExampleResponseDTO, errors.MessageErr) {
	return nil, nil
}

func (s *exampleService) DeleteExample(id uuid.UUID, ctx context.Context) (*dto.DeleteExampleResponseDTO, errors.MessageErr) {
	return nil, nil
}

func NewExampleService(repo example_repository.ExampleRepository) ExampleService {
	return &exampleService{
		repo: repo,
	}
}
