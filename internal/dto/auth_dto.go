package dto

type LoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
} // @name LoginRequest

type RegisterRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"required,oneof=TELLER CUSTOMER"`
} // @name RegisterRequest

type LoginResponseDTO struct {
	BaseResponse
	AccessToken string `json:"access_token"`
} // @name LoginResponse

type RegisterResponseDTO struct {
	BaseResponse
} // @name RegisterResponse
