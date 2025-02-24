package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/irfan44/go-http-boilerplate/docs"
	example_handler "github.com/irfan44/go-http-boilerplate/internal/domain/example/handler"
	example_service "github.com/irfan44/go-http-boilerplate/internal/domain/example/service"
	example_repo "github.com/irfan44/go-http-boilerplate/internal/repository/example"
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
	}

	services struct {
		exampleService example_service.ExampleService
	}
)

func (s *server) initializeRepositories() *repositories {
	exampleRepo := example_repo.NewExampleRepository(s.db)

	return &repositories{
		exampleRepository: exampleRepo,
	}
}

func (s *server) initializeServices(repo *repositories) *services {
	exampleService := example_service.NewExampleService(repo.exampleRepository)

	return &services{
		exampleService: exampleService,
	}
}

func (s *server) initializeHandlers(svc *services, v *validator.Validate, ctx context.Context) {
	exampleHandler := example_handler.NewExampleHandler(svc.exampleService, s.r, v, ctx)
	exampleHandler.MapRoutes()
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
	query := ``

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

	repo := s.initializeRepositories()
	svc := s.initializeServices(repo)
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
