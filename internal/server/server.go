package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/irfan44/go-http-boilerplate/internal/domain/example/handler"
	"github.com/irfan44/go-http-boilerplate/internal/domain/example/service"
	"github.com/irfan44/go-http-boilerplate/internal/repository/example"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
	"github.com/irfan44/go-http-boilerplate/internal/config"
)

type (
	server struct {
		cfg config.Config
		mux *http.ServeMux
		db  *sql.DB
	}

	repositories struct {
		exampleRepository repository.ExampleRepository
	}

	services struct {
		exampleService service.ExampleService
	}
)

func (s *server) initializeRepositories() *repositories {
	exampleRepo := repository.NewExampleRepository(s.db)

	return &repositories{
		exampleRepository: exampleRepo,
	}
}

func (s *server) initializeServices(repo *repositories) *services {
	exampleService := service.NewExampleService(repo.exampleRepository)

	return &services{
		exampleService: exampleService,
	}
}

func (s *server) initializeHandlers(svc *services, v *validator.Validate, ctx context.Context) {
	exampleHandler := handler.NewExampleHandler(svc.exampleService, s.mux, v, ctx)
	exampleHandler.MapRoutes()
}

func (s *server) initializeServer() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("Server listening on PORT %s\n", s.cfg.Http.Port)
		if err := http.ListenAndServe(s.cfg.Http.Port, s.mux); err != nil {
			log.Printf("Server error: %s\n", err.Error())
		}
	}()

	osCall := <-ch

	fmt.Printf("Server shutdown: %+v\n", osCall)
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

	s.initializeServer()
}

func NewServer(cfg config.Config, db *sql.DB) *server {
	return &server{
		cfg: cfg,
		mux: http.NewServeMux(),
		db:  db,
	}
}
