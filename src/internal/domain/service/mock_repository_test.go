package service

import (
	"context"
	"errors"
	"sync"

	"project03/internal/domain/models"
)

type mockGameRepository struct {
	mu    sync.RWMutex
	games map[string]models.Game
}

func newMockGameRepository() *mockGameRepository {
	return &mockGameRepository{
		games: make(map[string]models.Game),
	}
}

func (r *mockGameRepository) Save(ctx context.Context, game models.Game) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.games[game.ID] = game
	return nil
}

func (r *mockGameRepository) Get(ctx context.Context, id string) (models.Game, error) {
	if err := ctx.Err(); err != nil {
		return models.Game{}, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	game, ok := r.games[id]
	if !ok {
		return models.Game{}, errors.New("game not found")
	}

	return game, nil
}

func (r *mockGameRepository) Delete(ctx context.Context, id string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.games, id)
	return nil
}
