package server

import (
	"github.com/irfan44/go-http-boilerplate/internal/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) runGinServer() error {
	s.r.Use(middleware.ApplicationJsonResponseMiddleware())
	s.r.Use(middleware.EnableCorsMiddleware())

	s.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return s.r.Run(s.cfg.Http.Port)
}
