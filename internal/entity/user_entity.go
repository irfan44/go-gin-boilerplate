package entity

import (
	"github.com/golang-jwt/jwt"
	"github.com/irfan44/go-gin-boilerplate/internal/dto"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *User) ToUserDTO() *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		Id:        e.Id,
		Username:  e.Username,
		Role:      e.Role,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

type Users []User

func (e Users) ToUsersDTO() []dto.UserResponseDTO {
	result := []dto.UserResponseDTO{}

	for _, example := range e {
		result = append(result, *example.ToUserDTO())
	}

	return result
}

func (e *User) HashPassword() errs.MessageErr {
	b, err := bcrypt.GenerateFromPassword([]byte(e.Password), 8)

	if err != nil {
		return errs.NewInternalServerError()
	}

	e.Password = string(b)

	return nil
}

func (e *User) Compare(password string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(e.Password),
		[]byte(password),
	); err != nil {
		return errs.NewUnauthorizedError("invalid password")
	}

	return nil
}

func (e *User) NewClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"id":       e.Id,
		"username": e.Username,
		"role":     e.Role,
		"expr":     time.Now().Add(24 * time.Hour),
	}
}
