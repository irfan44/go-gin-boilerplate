package dto

import (
	"time"
)

type UserRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=TELLER CUSTOMER ADMIN"`
} // @name UserRequest

type UpdateUserRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=TELLER CUSTOMER ADMIN"`
} // @name UpdateUserRequest

type UserResponseDTO struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name UserResponse

type GetUsersResponseDTO struct {
	BaseResponse
	Data []UserResponseDTO `json:"data"`
} // @name GetUsersResponse

type GetUserByIdResponseDTO struct {
	BaseResponse
	Data UserResponseDTO `json:"data"`
} // @name GetUserByIdResponse

type CreateUserResponseDTO struct {
	BaseResponse
	Data UserResponseDTO `json:"data"`
} // @name CreateUserResponse

type UpdateUserResponseDTO struct {
	BaseResponse
	Data UserResponseDTO `json:"data"`
} // @name UpdateUserResponse

type DeleteUserResponseDTO struct {
	BaseResponse
} // @name DeleteUserResponse
