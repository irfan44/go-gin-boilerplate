package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"
)

func applicationJsonResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !strings.Contains(c.Request.URL.Path, "/swagger/") {
			c.Writer.Header().Set("Content-Type", "application/json")
		}

		c.Next()
	}
}

func enableCorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE"},
	})
}

func (s *server) runGinServer() error {
	s.r.Use(applicationJsonResponseMiddleware())
	s.r.Use(enableCorsMiddleware())
	s.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return s.r.Run(s.cfg.Http.Port)
}
