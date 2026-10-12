package service

import (
	"context"

	"project03/internal/domain/models"
)

type UserService interface {
	Register(ctx context.Context, login, password string) (models.User, error)
	GetByID(ctx context.Context, id string) (models.User, error)
}
