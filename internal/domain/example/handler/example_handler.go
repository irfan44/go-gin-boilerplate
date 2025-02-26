package example_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/irfan44/go-gin-boilerplate/internal/domain/example/service"
	"github.com/irfan44/go-gin-boilerplate/internal/dto"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"net/http"
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
	ctx := c.Request.Context()
	result, err := h.svc.GetExamples(ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result.Data)
}

// @Summary Get Example by ID
// @Tags examples
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} GetExampleByIdResponse
// @Router /examples/{id} [get]
func (h *exampleHandler) GetExampleById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := uuid.Parse(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	result, err := h.svc.GetExampleById(ctx, parsedId)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result.Data)
}

// @Summary Create Example
// @Tags examples
// @Accept json
// @Produce json
// @Param requestBody body ExampleRequest true "Request Body"
// @Success 200 {object} CreateExampleResponse
// @Router /examples [post]
func (h *exampleHandler) CreateExample(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.ExampleRequestDTO{}

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := h.svc.CreateExample(ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
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
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := uuid.Parse(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	payload := dto.ExampleRequestDTO{}

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	result, err := h.svc.UpdateExample(ctx, parsedId, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete Example
// @Tags examples
// @Produce json
// @Param id path string true "Example ID"
// @Success 200 {object} UpdateExampleResponse
// @Router /examples/{id} [delete]
func (h *exampleHandler) DeleteExample(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := uuid.Parse(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	result, errData := h.svc.DeleteExample(ctx, parsedId)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusNoContent, result)
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
