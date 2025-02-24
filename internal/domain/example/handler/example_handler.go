package example_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-http-boilerplate/internal/domain/example/service"
)

type exampleHandler struct {
	svc example_service.ExampleService
	r   *gin.Engine
	v   *validator.Validate
	ctx context.Context
}

// TODO: 5. adjust handler

func (h *exampleHandler) GetExamples(c *gin.Context) {

}

func (h *exampleHandler) GetExampleById(c *gin.Context) {

}

func (h *exampleHandler) CreateExample(c *gin.Context) {

}

func (h *exampleHandler) UpdateExample(c *gin.Context) {

}

func (h *exampleHandler) DeleteExample(c *gin.Context) {

}

func NewExampleHandler(
	svc example_service.ExampleService,
	r *gin.Engine,
	v *validator.Validate,
	ctx context.Context,
) *exampleHandler {
	return &exampleHandler{
		svc: svc,
		r:   r,
		v:   v,
		ctx: ctx,
	}
}
