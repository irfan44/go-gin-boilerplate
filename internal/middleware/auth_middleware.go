package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/irfan44/go-gin-boilerplate/internal/config"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"github.com/irfan44/go-gin-boilerplate/pkg/internal_jwt"
	"strings"
)

type (
	AuthMiddleware interface {
		Authentication() gin.HandlerFunc
		AdminAuthorization() gin.HandlerFunc
		TellerAuthorization() gin.HandlerFunc
	}

	authMiddleware struct {
		ctx         context.Context
		internalJwt internal_jwt.InternalJwt
		cfg         config.Config
	}
)

func (m *authMiddleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		if publicPath := strings.HasPrefix(c.FullPath(), "/auth") || strings.HasPrefix(c.FullPath(), "/swagger"); publicPath {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")

		mapClaims, err := m.internalJwt.ValidateBearerToken(authHeader, m.cfg.Jwt.SecretKey)

		if err != nil {
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		id, ok := mapClaims["id"].(float64)

		if !ok {
			err = errs.NewUnauthenticatedError("Invalid token.")
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		username, ok := mapClaims["username"].(string)

		if !ok {
			err = errs.NewUnauthenticatedError("Invalid token.")
			c.JSON(err.StatusCode(), err)
			return
		}

		role, ok := mapClaims["role"].(string)

		if !ok {
			err = errs.NewUnauthenticatedError("Invalid token.")
			c.JSON(err.StatusCode(), err)
			return
		}

		c.Set("userId", int(id))
		c.Set("username", username)
		c.Set("role", role)

		c.Next()
	}
}

func (m *authMiddleware) AdminAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Value("role").(string)

		if !ok || role != "ADMIN" {
			err := errs.NewUnauthorizedError("Cannot access endpoint.")
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		c.Next()
	}
}

func (m *authMiddleware) TellerAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Value("role").(string)

		if !ok || role != "TELLER" {
			err := errs.NewUnauthorizedError("Cannot access endpoint.")
			c.AbortWithStatusJSON(err.StatusCode(), err)
			return
		}

		c.Next()
	}
}

func NewAuthMiddleware(
	internalJwt internal_jwt.InternalJwt,
	cfg config.Config,
) AuthMiddleware {
	return &authMiddleware{
		internalJwt: internalJwt,
		cfg:         cfg,
	}
}
