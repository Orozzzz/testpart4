package repository

import (
    "errors"
    "sync"
    
    "project03/internal/domain/models"
    "project03/internal/datasource/mappers"
    dsmodels "project03/internal/datasource/models"
)

type gameRepository struct {
	storage *sync.Map 
}

func NewGameRepository() GameRepository {
	return &gameRepository{
		storage: &sync.Map{},
	}
}

func (r *gameRepository) Save(game models.Game) error {
	dto := mappers.ToDTO(game)
	r.storage.Store(game.ID, dto)
	return nil
}

func (r *gameRepository) Get(id string) (models.Game, error) {
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

func (r *gameRepository) Delete(id string) error {
	r.storage.Delete(id)
	return nil
}
