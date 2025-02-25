package example_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/irfan44/go-http-boilerplate/internal/repository/example"
	"github.com/irfan44/go-http-boilerplate/pkg/errs"

	"github.com/irfan44/go-http-boilerplate/internal/dto"
)

type (
	ExampleService interface {
		GetExamples(ctx context.Context) (*dto.GetExamplesResponseDTO, errs.MessageErr)
		GetExampleById(ctx context.Context, id uuid.UUID) (*dto.GetExampleByIdResponseDTO, errs.MessageErr)
		CreateExample(ctx context.Context, example dto.ExampleRequestDTO) (*dto.CreateExampleResponseDTO, errs.MessageErr)
		UpdateExample(ctx context.Context, id uuid.UUID, example dto.ExampleRequestDTO) (*dto.UpdateExampleResponseDTO, errs.MessageErr)
		DeleteExample(ctx context.Context, id uuid.UUID) (*dto.DeleteExampleResponseDTO, errs.MessageErr)
	}

	exampleService struct {
		repo example_repository.ExampleRepository
	}
)

// TODO: 4. adjust service

func (s *exampleService) GetExamples(ctx context.Context) (*dto.GetExamplesResponseDTO, errs.MessageErr) {
	return nil, nil
}

func (s *exampleService) GetExampleById(ctx context.Context, id uuid.UUID) (*dto.GetExampleByIdResponseDTO, errs.MessageErr) {
	return nil, nil
}

func (s *exampleService) CreateExample(ctx context.Context, example dto.ExampleRequestDTO) (*dto.CreateExampleResponseDTO, errs.MessageErr) {
	return nil, nil
}

func (s *exampleService) UpdateExample(ctx context.Context, id uuid.UUID, example dto.ExampleRequestDTO) (*dto.UpdateExampleResponseDTO, errs.MessageErr) {
	return nil, nil
}

func (s *exampleService) DeleteExample(ctx context.Context, id uuid.UUID) (*dto.DeleteExampleResponseDTO, errs.MessageErr) {
	return nil, nil
}

func NewExampleService(repo example_repository.ExampleRepository) ExampleService {
	return &exampleService{
		repo: repo,
	}
}
