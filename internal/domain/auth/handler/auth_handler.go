package auth_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-gin-boilerplate/internal/domain/auth/service"
	"github.com/irfan44/go-gin-boilerplate/internal/dto"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"net/http"
)

type authHandler struct {
	svc auth_service.AuthService
	r   *gin.Engine
	v   *validator.Validate
	ctx context.Context
}

// TODO: 5. adjust handler

// @Summary Login
// @Tags auth
// @Accept json
// @Produce json
// @Param requestBody body LoginRequest true "Request Body"
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func (h *authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.LoginRequestDTO{}

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	if errVal := h.v.Struct(payload); errVal != nil {
		errMsg := errs.NewBadRequest(errVal.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	result, err := h.svc.Login(ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Register
// @Tags auth
// @Accept json
// @Produce json
// @Param requestBody body RegisterRequest true "Request Body"
// @Success 200 {object} RegisterResponse
// @Router /auth/register [post]
func (h *authHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.RegisterRequestDTO{}

	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		errData := errs.NewUnprocessibleEntityError(err.Error())
		c.JSON(errData.StatusCode(), errData)
		return
	}

	if errVal := h.v.Struct(payload); errVal != nil {
		errMsg := errs.NewBadRequest(errVal.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
		return
	}

	result, err := h.svc.Register(ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func NewExampleHandler(
	svc auth_service.AuthService,
	r *gin.Engine,
	v *validator.Validate,
	ctx context.Context,
) *authHandler {
	return &authHandler{
		svc: svc,
		r:   r,
		v:   v,
		ctx: ctx,
	}
}
