package example_repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/irfan44/go-http-boilerplate/pkg/errors"

	"github.com/irfan44/go-http-boilerplate/internal/entity"
)

type (
	ExampleRepository interface {
		GetExamples(ctx context.Context) (entity.Examples, errors.MessageErr)
		GetExampleById(id uuid.UUID, ctx context.Context) (*entity.Example, errors.MessageErr)
		GetExampleByName(name string, ctx context.Context) (*entity.Example, errors.MessageErr)
		CreateExample(example entity.Example, ctx context.Context) (*entity.Example, errors.MessageErr)
		UpdateExample(example entity.Example, id uuid.UUID, ctx context.Context) (*entity.Example, errors.MessageErr)
		DeleteExample(id uuid.UUID, ctx context.Context) (bool, errors.MessageErr)
	}

	exampleRepository struct {
		db *sql.DB
	}
)

// TODO: 3. adjust repository

func (r *exampleRepository) GetExamples(ctx context.Context) (entity.Examples, errors.MessageErr) {
	return nil, nil
}

func (r *exampleRepository) GetExampleById(id uuid.UUID, ctx context.Context) (*entity.Example, errors.MessageErr) {
	return nil, nil
}

func (r *exampleRepository) GetExampleByName(name string, ctx context.Context) (*entity.Example, errors.MessageErr) {
	return nil, nil
}

func (r *exampleRepository) CreateExample(example entity.Example, ctx context.Context) (*entity.Example, errors.MessageErr) {
	return nil, nil
}

func (r *exampleRepository) UpdateExample(example entity.Example, id uuid.UUID, ctx context.Context) (*entity.Example, errors.MessageErr) {
	return nil, nil
}

func (r *exampleRepository) DeleteExample(id uuid.UUID, ctx context.Context) (bool, errors.MessageErr) {
	return false, nil
}

func NewExampleRepository(db *sql.DB) ExampleRepository {
	return &exampleRepository{
		db: db,
	}
}
