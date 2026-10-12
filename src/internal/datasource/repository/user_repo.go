package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"project03/internal/datasource/dbmodels"
	"project03/internal/datasource/mappers"
	"project03/internal/domain/models"
	"project03/internal/domain/ports"
)

type userRepository struct {
	pool *pgxpool.Pool
}

var _ ports.UserRepository = (*userRepository)(nil)

func NewUserRepository(pool *pgxpool.Pool) ports.UserRepository {
	return &userRepository{pool: pool}
}

func (r *userRepository) Create(ctx context.Context, user models.User) error {
	dto := mappers.UserToDTO(user)

	if dto.ID == uuid.Nil {
		return errors.New("invalid user ID")
	}

	_, err := r.pool.Exec(ctx, `
		INSERT INTO users (id, login, password_hash)
		VALUES ($1, $2, $3)
	`, dto.ID, dto.Login, dto.PasswordHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) &&
			pgErr.Code == "23505" && //код PostgreSQL для нарушения уникальности
			pgErr.ConstraintName == "users_login_key" {
			return models.ErrLoginAlreadyTaken
		}

		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (r *userRepository) GetByLogin(
	ctx context.Context,
	login string,
) (models.User, error) {
	var dto dbmodels.UserDTO

	err := r.pool.QueryRow(ctx, `
		SELECT id, login, password_hash
		FROM users
		WHERE login = $1
	`, login).Scan(
		&dto.ID,
		&dto.Login,
		&dto.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}

		return models.User{}, fmt.Errorf("get user by login: %w", err)
	}

	return mappers.UserToDomain(dto), nil
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id string,
) (models.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return models.User{}, fmt.Errorf("invalid user ID: %w", err)
	}

	var dto dbmodels.UserDTO

	err = r.pool.QueryRow(ctx, `
		SELECT id, login, password_hash
		FROM users
		WHERE id = $1
	`, userID).Scan(
		&dto.ID,
		&dto.Login,
		&dto.PasswordHash,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.User{}, models.ErrUserNotFound
		}

		return models.User{}, fmt.Errorf("get user by ID: %w", err)
	}

	return mappers.UserToDomain(dto), nil
}
