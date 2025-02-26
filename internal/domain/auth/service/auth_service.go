package auth_service

import (
	"context"
	"github.com/irfan44/go-gin-boilerplate/internal/config"
	"github.com/irfan44/go-gin-boilerplate/internal/entity"
	"github.com/irfan44/go-gin-boilerplate/internal/repository/user"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"github.com/irfan44/go-gin-boilerplate/pkg/internal_http"
	"github.com/irfan44/go-gin-boilerplate/pkg/internal_jwt"
	"time"

	"github.com/irfan44/go-gin-boilerplate/internal/dto"
)

type (
	AuthService interface {
		Login(ctx context.Context, user dto.LoginRequestDTO) (*dto.LoginResponseDTO, errs.MessageErr)
		Register(ctx context.Context, user dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, errs.MessageErr)
	}

	authService struct {
		repo        user_repository.UserRepository
		internalJwt internal_jwt.InternalJwt
		cfg         config.Config
	}
)

// TODO: 4. adjust service

func (s *authService) Login(ctx context.Context, auth dto.LoginRequestDTO) (*dto.LoginResponseDTO, errs.MessageErr) {
	res, err := s.repo.GetUserByName(ctx, auth.Username)

	if err != nil {
		return nil, err
	}

	if err = res.Compare(auth.Password); err != nil {
		return nil, errs.NewUnauthorizedError("Incorrect username/password")
	}

	claim := res.NewClaim()

	token := s.internalJwt.GenerateToken(claim, s.cfg.Jwt.SecretKey)

	result := dto.LoginResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("Login successfully"),
		AccessToken:  token,
	}

	return &result, nil
}

func (s *authService) Register(ctx context.Context, auth dto.RegisterRequestDTO) (*dto.RegisterResponseDTO, errs.MessageErr) {
	resCheck, errCheck := s.repo.GetUserByName(ctx, auth.Username)

	if errCheck != nil && errCheck.StatusCode() != 404 {
		return nil, errCheck
	}

	if resCheck != nil {
		return nil, errs.NewConflictError("Please use another username")
	}

	payload := entity.User{
		Id:        0,
		Username:  auth.Username,
		Password:  auth.Password,
		Role:      auth.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if errHash := payload.HashPassword(); errHash != nil {
		return nil, errHash
	}

	_, err := s.repo.CreateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	result := dto.RegisterResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("Registered successfully"),
	}

	return &result, nil
}

func NewAuthService(repo user_repository.UserRepository, internalJwt internal_jwt.InternalJwt, cfg config.Config) AuthService {
	return &authService{
		repo:        repo,
		internalJwt: internalJwt,
		cfg:         cfg,
	}
}
