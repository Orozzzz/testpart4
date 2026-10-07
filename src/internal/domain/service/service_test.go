package service

import (
	"testing"

	"project03/internal/datasource/repository"
	"project03/internal/domain/models"
)

func setupTest() GameService {
	repo := repository.NewGameRepository()
	return NewService(repo)
}

func TestNewGame(t *testing.T) {
	service := setupTest()
	game := service.NewGame()

	if game.ID == "" {
		t.Error("Expected game ID to be set")
	}
}

func TestSaveAndGetGame(t *testing.T) {
	service := setupTest()

	game := service.NewGame()

	game.Board[0][0] = 1

	err := service.SaveGame(game)
	if err != nil {
		t.Errorf("Error saving game: %v", err)
	}

	loadedGame, err := service.GetGame(game.ID)
	if err != nil {
		t.Errorf("Error getting game: %v", err)
	}

	if loadedGame.Board[0][0] != 1 {
		t.Errorf("Expected board[0][0] = 1, got %v", loadedGame.Board[0][0])
	}
}

func TestMinimax(t *testing.T) {
	service := setupTest()

	board := models.Board{
		{1, 1, 0},
		{0, 2, 0},
		{0, 0, 2},
	}

	game := models.Game{
		ID:    "test",
		Board: board,
		Turn:  1,
	}

	newBoard, err := service.GetNextMove(game)
	if err != nil {
		t.Errorf("Error getting next move: %v", err)
	}

	if newBoard[0][2] != 2 {
		t.Errorf("Expected compter to play at (0,2), got %v", newBoard)
	}
}
