package server

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) runGinServer() error {
	s.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return s.r.Run(s.cfg.Http.Port)
}
