package user_handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-http-boilerplate/internal/domain/user/service"
	"github.com/irfan44/go-http-boilerplate/internal/dto"
	"github.com/irfan44/go-http-boilerplate/internal/middleware"
	"github.com/irfan44/go-http-boilerplate/pkg/errs"
	"net/http"
	"strconv"
)

type userHandler struct {
	svc user_service.UserService
	r   *gin.Engine
	v   *validator.Validate
	ctx context.Context
	m   middleware.AuthMiddleware
}

// @Summary Get All Users
// @Tags users
// @Produce json
// @Success 200 {object} GetUsersResponse
// @Router /users [get]
func (h *userHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := h.svc.GetUsers(ctx)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Get User by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} GetUserByIdResponse
// @Router /users/{id} [get]
func (h *userHandler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := strconv.Atoi(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	result, err := h.svc.GetUserById(ctx, parsedId)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create User
// @Tags users
// @Accept json
// @Produce json
// @Param requestBody body UserRequest true "Request Body"
// @Success 200 {object} CreateUserResponse
// @Router /users [post]
func (h *userHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	payload := dto.UserRequestDTO{}

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

	result, err := h.svc.CreateUser(ctx, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// @Summary Update User
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param requestBody body UpdateUserRequest true "Request Body"
// @Success 200 {object} UpdateUserResponse
// @Router /users/{id} [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := strconv.Atoi(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	payload := dto.UpdateUserRequestDTO{}

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

	result, err := h.svc.UpdateUser(ctx, parsedId, payload)

	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Delete User
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UpdateUserResponse
// @Router /users/{id} [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	parsedId, errParse := strconv.Atoi(id)

	if errParse != nil {
		errMsg := errs.NewBadRequest(errParse.Error())
		c.JSON(errMsg.StatusCode(), errMsg)
	}

	result, errData := h.svc.DeleteUser(ctx, parsedId)

	if errData != nil {
		c.JSON(errData.StatusCode(), errData)
		return
	}

	c.JSON(http.StatusOK, result)
}

func NewUserHandler(
	svc user_service.UserService,
	r *gin.Engine,
	v *validator.Validate,
	ctx context.Context,
	m middleware.AuthMiddleware,
) *userHandler {
	return &userHandler{
		svc: svc,
		r:   r,
		v:   v,
		ctx: ctx,
		m:   m,
	}
}
