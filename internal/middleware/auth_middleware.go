package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/irfan44/go-http-boilerplate/internal/config"
	"github.com/irfan44/go-http-boilerplate/internal/domain/user/service"
	"github.com/irfan44/go-http-boilerplate/pkg/errs"
	"github.com/irfan44/go-http-boilerplate/pkg/internal_jwt"
	"strings"
)

type (
	AuthMiddleware interface {
		Authentication(c *gin.Context)
	}

	authMiddleware struct {
		ctx         context.Context
		internalJwt internal_jwt.InternalJwt
		cfg         config.Config
		userService user_service.UserService
	}
)

func (m *authMiddleware) Authentication(c *gin.Context) {
	path := strings.HasPrefix(c.FullPath(), "/auth") || strings.HasPrefix(c.FullPath(), "/swagger")

	if path {
		c.Next()
		return
	}

	authHeader, ok := c.Get("Authorization")

	if !ok {
		err := errs.NewUnauthenticatedError("Invalid token.")
		c.JSON(err.StatusCode(), err)
	}

	token := authHeader.(string)

	mapClaims, err := m.internalJwt.ValidateBearerToken(token, m.cfg.Jwt.SecretKey)

	if err != nil {
		c.JSON(err.StatusCode(), err)
	}

	id, ok := mapClaims["id"].(float64)

	if !ok {
		err = errs.NewUnauthenticatedError("Invalid token.")
		c.JSON(err.StatusCode(), err)
	}

	role, ok := mapClaims["role"].(string)

	if !ok {
		err = errs.NewUnauthenticatedError("Invalid token.")
		c.JSON(err.StatusCode(), err)
	}

	ctx := c.Request.Context()

	if _, err = m.userService.GetUserById(ctx, int(id)); err != nil {
		c.JSON(err.StatusCode(), err)
	}

	c.Set("userId", int(id))
	c.Set("role", role)

	c.Next()
}

func NewAuthMiddleware(
	internalJwt internal_jwt.InternalJwt,
	cfg config.Config,
	userService user_service.UserService,
) AuthMiddleware {
	return &authMiddleware{
		internalJwt: internalJwt,
		cfg:         cfg,
		userService: userService,
	}
}
