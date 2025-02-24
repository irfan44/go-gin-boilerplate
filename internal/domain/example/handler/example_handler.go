package handler

import (
	"context"
	"github.com/irfan44/go-http-boilerplate/internal/domain/example/service"
	"github.com/irfan44/go-http-boilerplate/pkg/internal_http"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type exampleHandler struct {
	svc service.ExampleService
	mux *http.ServeMux
	v   *validator.Validate
	ctx context.Context
}

// TODO: 5. adjust handler

func (h *exampleHandler) GetExamples() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal_http.ApplyJSON(w)
	}
}

func (h *exampleHandler) GetExampleById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal_http.ApplyJSON(w)
	}
}

func (h *exampleHandler) CreateExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal_http.ApplyJSON(w)
	}
}

func (h *exampleHandler) UpdateExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal_http.ApplyJSON(w)
	}
}

func (h *exampleHandler) DeleteExample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		internal_http.ApplyJSON(w)
	}
}

func NewExampleHandler(
	svc service.ExampleService,
	mux *http.ServeMux,
	v *validator.Validate,
	ctx context.Context,
) *exampleHandler {
	return &exampleHandler{
		svc: svc,
		mux: mux,
		v:   v,
		ctx: ctx,
	}
}
