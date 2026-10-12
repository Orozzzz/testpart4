package models

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Login        string
	PasswordHash string
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrLoginAlreadyTaken = errors.New("login already taken")
)
