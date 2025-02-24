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

// @Summary Get All Examples
// @Tags examples
// @Produce json
// @Success 200 {object} GetExamplesResponse
// @Router /examples [get]
func (h *exampleHandler) GetExamples(c *gin.Context) {

}

// @Summary Get Example by ID
// @Tags examples
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} GetExampleByIdResponse
// @Router /examples/{id} [get]
func (h *exampleHandler) GetExampleById(c *gin.Context) {

}

// @Summary Create Example
// @Tags examples
// @Accept json
// @Produce json
// @Param requestBody body ExampleRequest true "Request Body"
// @Success 200 {object} CreateExampleResponse
// @Router /examples [post]
func (h *exampleHandler) CreateExample(c *gin.Context) {

}

// @Summary Update Example
// @Tags examples
// @Accept json
// @Produce json
// @Param id path string true "Example ID"
// @Param requestBody body ExampleRequest true "Request Body"
// @Success 200 {object} UpdateExampleResponse
// @Router /examples/{id} [put]
func (h *exampleHandler) UpdateExample(c *gin.Context) {

}

// @Summary Delete Example
// @Tags examples
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} UpdateExampleResponse
// @Router /examples/{id} [delete]
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
