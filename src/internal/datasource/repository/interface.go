package repository

import "project03/internal/domain/models"

type GameRepository interface {
    Save(game models.Game) error
    Get(id string) (models.Game, error)
    Delete(id string) error
}