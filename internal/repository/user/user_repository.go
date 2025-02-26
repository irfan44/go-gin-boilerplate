package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/irfan44/go-gin-boilerplate/pkg/errs"
	"log"

	"github.com/irfan44/go-gin-boilerplate/internal/entity"
)

type (
	UserRepository interface {
		GetUsers(ctx context.Context) (entity.Users, errs.MessageErr)
		GetUserById(ctx context.Context, id int) (*entity.User, errs.MessageErr)
		GetUserByName(ctx context.Context, name string) (*entity.User, errs.MessageErr)
		CreateUser(ctx context.Context, user entity.User) (*entity.User, errs.MessageErr)
		UpdateUser(ctx context.Context, id int, user entity.User) (*entity.User, errs.MessageErr)
		DeleteUser(ctx context.Context, id int) (bool, errs.MessageErr)
	}

	userRepository struct {
		db *sql.DB
	}
)

func (r *userRepository) GetUsers(ctx context.Context) (entity.Users, errs.MessageErr) {
	rows, err := r.db.QueryContext(ctx, GET_USERS)

	if err != nil {
		log.Printf("db scan get all users: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	results := entity.Users{}

	for rows.Next() {
		user := entity.User{}

		if err = rows.Scan(
			&user.Id,
			&user.Username,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			log.Printf("Error get all users: %s\n", err.Error())
			return nil, errs.NewInternalServerError()
		}

		results = append(results, user)
	}

	return results, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := r.db.QueryRowContext(ctx, GET_USER_BY_ID, id).Scan(
		&user.Id,
		&user.Username,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		log.Printf("Error get user by id: %s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError(fmt.Sprintf("User with id %d not found", id))
		}

		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}

func (r *userRepository) GetUserByName(ctx context.Context, name string) (*entity.User, errs.MessageErr) {
	user := entity.User{}

	if err := r.db.QueryRowContext(ctx, GET_USER_BY_NAME, name).Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		log.Printf("Error get user by name: %s\n", err.Error())

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError(fmt.Sprintf("User with username %s not found", name))
		}

		return nil, errs.NewInternalServerError()
	}

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (*entity.User, errs.MessageErr) {
	newUser := entity.User{}

	if err := r.db.QueryRowContext(ctx, CREATE_USER, user.Username, user.Password, user.Role).Scan(
		&newUser.Id,
		&newUser.Username,
		&newUser.Role,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	); err != nil {
		log.Printf("Error create user: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	return &newUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (*entity.User, errs.MessageErr) {
	updatedUser := entity.User{}

	if err := r.db.QueryRowContext(ctx, UPDATE_USER, user.Username, user.Role, user.UpdatedAt, id).Scan(
		&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.Role,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	); err != nil {
		log.Printf("Error update user: %s\n", err.Error())
		return nil, errs.NewInternalServerError()
	}

	return &updatedUser, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) (bool, errs.MessageErr) {
	_, err := r.db.ExecContext(ctx, DELETE_USER, id)

	if err != nil {
		log.Printf("Error delete user: %s", err.Error())
		return false, errs.NewInternalServerError()
	}

	return true, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
