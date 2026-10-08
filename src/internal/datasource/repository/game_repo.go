package repository

import (
	"context"
	"errors"
	"sync"

	"project03/internal/datasource/mappers"
	dsmodels "project03/internal/datasource/models"
	"project03/internal/domain/models"
	"project03/internal/domain/ports"
)

type gameRepository struct {
	storage *sync.Map
}

var _ ports.GameRepository = (*gameRepository)(nil)

func NewGameRepository() ports.GameRepository {
	return &gameRepository{
		storage: &sync.Map{},
	}
}

func (r *gameRepository) Save(ctx context.Context, game models.Game) error {
	_ = ctx

	dto := mappers.ToDTO(game)
	r.storage.Store(game.ID, dto)

	return nil
}

func (r *gameRepository) Get(ctx context.Context, id string) (models.Game, error) {
	_ = ctx

	val, ok := r.storage.Load(id)
	if !ok {
		return models.Game{}, errors.New("game not found")
	}

	dto, ok := val.(dsmodels.GameDTO)
	if !ok {
		return models.Game{}, errors.New("invalid data in storage")
	}

	return mappers.ToDomain(dto), nil
}

func (r *gameRepository) Delete(ctx context.Context, id string) error {
	_ = ctx

	r.storage.Delete(id)

	return nil
}
