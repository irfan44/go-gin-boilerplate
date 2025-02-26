package user_service

import (
	"context"
	"github.com/irfan44/go-gin-boilerplate/internal/entity"
	"github.com/irfan44/go-gin-boilerplate/internal/repository/user"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"github.com/irfan44/go-gin-boilerplate/pkg/internal_http"
	"time"

	"github.com/irfan44/go-gin-boilerplate/internal/dto"
)

type (
	UserService interface {
		GetUsers(ctx context.Context) (*dto.GetUsersResponseDTO, errs.MessageErr)
		GetUserById(ctx context.Context, id int) (*dto.GetUserByIdResponseDTO, errs.MessageErr)
		CreateUser(ctx context.Context, user dto.UserRequestDTO) (*dto.CreateUserResponseDTO, errs.MessageErr)
		UpdateUser(ctx context.Context, id int, user dto.UpdateUserRequestDTO) (*dto.UpdateUserResponseDTO, errs.MessageErr)
		DeleteUser(ctx context.Context, id int) (*dto.DeleteUserResponseDTO, errs.MessageErr)
	}

	userService struct {
		repo user_repository.UserRepository
	}
)

// TODO: 4. adjust service

func (s *userService) GetUsers(ctx context.Context) (*dto.GetUsersResponseDTO, errs.MessageErr) {
	res, err := s.repo.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetUsersResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("Get users successfully"),
		Data:         res.ToUsersDTO(),
	}

	return &result, nil
}

func (s *userService) GetUserById(ctx context.Context, id int) (*dto.GetUserByIdResponseDTO, errs.MessageErr) {
	res, err := s.repo.GetUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	result := dto.GetUserByIdResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("Get user successfully"),
		Data:         *res.ToUserDTO(),
	}

	return &result, nil
}

func (s *userService) CreateUser(ctx context.Context, user dto.UserRequestDTO) (*dto.CreateUserResponseDTO, errs.MessageErr) {
	resCheck, errCheck := s.repo.GetUserByName(ctx, user.Username)

	if errCheck != nil && errCheck.StatusCode() != 404 {
		return nil, errCheck
	}

	if resCheck != nil {
		return nil, errs.NewConflictError("Please use another username")
	}

	payload := entity.User{
		Id:        0,
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if errHash := payload.HashPassword(); errHash != nil {
		return nil, errHash
	}

	res, err := s.repo.CreateUser(ctx, payload)

	if err != nil {
		return nil, err
	}

	result := dto.CreateUserResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("User created successfully"),
		Data:         *res.ToUserDTO(),
	}

	return &result, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int, user dto.UpdateUserRequestDTO) (*dto.UpdateUserResponseDTO, errs.MessageErr) {
	resCheckId, errCheckId := s.repo.GetUserById(ctx, id)

	if errCheckId != nil {
		return nil, errCheckId
	}

	resCheckName, errCheckName := s.repo.GetUserByName(ctx, user.Username)

	if resCheckName != nil && resCheckId.Username != user.Username {
		return nil, errs.NewConflictError("Please use another username")
	}

	if errCheckName != nil && errCheckName.StatusCode() != 404 {
		return nil, errCheckName
	}

	payload := entity.User{
		Id:        0,
		Username:  user.Username,
		Password:  "",
		Role:      user.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := s.repo.UpdateUser(ctx, id, payload)

	if err != nil {
		return nil, err
	}

	result := dto.UpdateUserResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("User updated successfully"),
		Data:         *res.ToUserDTO(),
	}

	return &result, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int) (*dto.DeleteUserResponseDTO, errs.MessageErr) {
	_, errCheck := s.repo.GetUserById(ctx, id)

	if errCheck != nil {
		return nil, errCheck
	}

	_, err := s.repo.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	result := dto.DeleteUserResponseDTO{
		BaseResponse: internal_http.NewBaseResponse("User deleted successfully"),
	}

	return &result, nil
}

func NewUserService(repo user_repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
