package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/irfan44/go-http-boilerplate/docs"
	auth_handler "github.com/irfan44/go-http-boilerplate/internal/domain/auth/handler"
	auth_service "github.com/irfan44/go-http-boilerplate/internal/domain/auth/service"
	example_handler "github.com/irfan44/go-http-boilerplate/internal/domain/example/handler"
	example_service "github.com/irfan44/go-http-boilerplate/internal/domain/example/service"
	user_handler "github.com/irfan44/go-http-boilerplate/internal/domain/user/handler"
	user_service "github.com/irfan44/go-http-boilerplate/internal/domain/user/service"
	"github.com/irfan44/go-http-boilerplate/internal/middleware"
	example_repo "github.com/irfan44/go-http-boilerplate/internal/repository/example"
	user_repo "github.com/irfan44/go-http-boilerplate/internal/repository/user"
	"github.com/irfan44/go-http-boilerplate/pkg/internal_jwt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-http-boilerplate/internal/config"
)

type (
	server struct {
		cfg config.Config
		r   *gin.Engine
		db  *sql.DB
	}

	repositories struct {
		exampleRepository example_repo.ExampleRepository
		userRepository    user_repo.UserRepository
	}

	services struct {
		exampleService example_service.ExampleService
		userService    user_service.UserService
		authService    auth_service.AuthService
	}
)

func (s *server) initializeRepositories() *repositories {
	exampleRepo := example_repo.NewExampleRepository(s.db)
	userRepo := user_repo.NewUserRepository(s.db)

	return &repositories{
		exampleRepository: exampleRepo,
		userRepository:    userRepo,
	}
}

func (s *server) initializeServices(repo *repositories, internalJwt internal_jwt.InternalJwt) *services {
	exampleService := example_service.NewExampleService(repo.exampleRepository)
	userService := user_service.NewUserService(repo.userRepository)
	authService := auth_service.NewAuthService(repo.userRepository, internalJwt, s.cfg)

	return &services{
		exampleService: exampleService,
		userService:    userService,
		authService:    authService,
	}
}

func (s *server) initializeHandlers(svc *services, v *validator.Validate, ctx context.Context) {
	exampleHandler := example_handler.NewExampleHandler(svc.exampleService, s.r, v, ctx)
	_ = exampleHandler
	//exampleHandler.MapRoutes()

	userHandler := user_handler.NewUserHandler(svc.userService, s.r, v, ctx)
	userHandler.MapRoutes()

	authHandler := auth_handler.NewExampleHandler(svc.authService, s.r, v, ctx)
	authHandler.MapRoutes()
}

func (s *server) initializeMiddleware(internalJwt internal_jwt.InternalJwt, svc *services) {
	s.r.Use(middleware.ApplicationJsonResponseMiddleware())
	s.r.Use(middleware.EnableCorsMiddleware())

	authMiddleware := middleware.NewAuthMiddleware(internalJwt, s.cfg, svc.userService)
	s.r.Use(authMiddleware.Authentication)
}

func (s *server) initializeServer() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on PORT %s\n", s.cfg.Http.Port)
		if err := s.runGinServer(); err != nil {
			log.Printf("Server error: %s\n", err.Error())
		}
	}()

	osCall := <-ch

	log.Printf("Server shutdown: %+v\n", osCall)
}

func (s *server) initializeTable() error {
	// TODO: fill init table query
	query := `
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL primary key,
		    username VARCHAR (255) UNIQUE NOT NULL,
		    password VARCHAR (255) NOT NULL,
		    role VARCHAR (30) NOT NULL CHECK (role IN ('TELLER', 'CUSTOMER', 'ADMIN')),
		    created_at TIMESTAMPTZ DEFAULT NOW(),
			updated_at TIMESTAMPTZ DEFAULT NOW()
		)
	`

	if _, err := s.db.Exec(query); err != nil {
		log.Printf("Initialize table error: %s\n", err.Error())

		if err = s.db.Close(); err != nil {
			log.Printf("Graceful DB shutdown: %s\n", err.Error())
		} else {
			log.Printf("Successfully graceful DB shutdown \n")
		}

		return err
	}

	log.Println("Successfully initiate table")

	return nil
}

func (s *server) initializeSwagger() {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", s.cfg.Http.Host, s.cfg.Http.Port)
}

func (s *server) Run() {
	if err := s.initializeTable(); err != nil {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	v := validator.New()

	internalJwt := internal_jwt.NewInternalJwt()

	repo := s.initializeRepositories()
	svc := s.initializeServices(repo, internalJwt)
	s.initializeMiddleware(internalJwt, svc)
	s.initializeHandlers(svc, v, ctx)

	s.initializeSwagger()

	s.initializeServer()
}

func NewServer(cfg config.Config, db *sql.DB) *server {
	return &server{
		cfg: cfg,
		r:   gin.Default(),
		db:  db,
	}
}
