package models

import (
	"testing"

	"github.com/google/uuid"
)

func TestGameMakeMove(t *testing.T) {
	playerID := uuid.New()
	opponentID := uuid.New()

	game := Game{
		ID:    "test-game",
		Board: NewBoard(),
		Players: []Player{
			{
				ID:   playerID,
				Cell: X,
			},
			{
				ID:   opponentID,
				Cell: O,
			},
		},
		CurrentPlayer: playerID,
		Status:        GameInProgress,
		Winner:        uuid.Nil,
	}

	err := game.MakeMove(playerID, 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if game.Board[0][0] != X {
		t.Fatalf("expected X, got %v", game.Board[0][0])
	}

	if game.CurrentPlayer != opponentID {
		t.Fatalf("expected next player %v, got %v", opponentID, game.CurrentPlayer)
	}

	if game.Status != GameInProgress {
		t.Fatalf("expected game to be in progress, got %v", game.Status)
	}
}

func TestGameMakeMoveRejectsOccupiedCell(t *testing.T) {
	playerID := uuid.New()

	game := Game{
		Board: NewBoard(),
		Players: []Player{
			{ID: playerID, Cell: X},
		},
		CurrentPlayer: playerID,
		Status:        GameInProgress,
	}

	game.Board[0][0] = X

	err := game.MakeMove(playerID, 0, 0)
	if err != ErrCellOccupied {
		t.Fatalf("expected ErrCellOccupied, got %v", err)
	}
}

func TestGameMakeMoveRejectsWrongPlayer(t *testing.T) {
	playerID := uuid.New()
	opponentID := uuid.New()

	game := Game{
		Board: NewBoard(),
		Players: []Player{
			{ID: playerID, Cell: X},
			{ID: opponentID, Cell: O},
		},
		CurrentPlayer: playerID,
		Status:        GameInProgress,
	}

	err := game.MakeMove(opponentID, 0, 0)
	if err != ErrNotPlayerTurn {
		t.Fatalf("expected ErrNotPlayerTurn, got %v", err)
	}
}

func TestGameMakeMoveRejectsUnknownPlayer(t *testing.T) {
	playerID := uuid.New()
	unknownID := uuid.New()

	game := Game{
		Board: NewBoard(),
		Players: []Player{
			{ID: playerID, Cell: X},
		},
		CurrentPlayer: playerID,
		Status:        GameInProgress,
	}

	err := game.MakeMove(unknownID, 0, 0)
	if err != ErrPlayerNotFound {
		t.Fatalf("expected ErrPlayerNotFound, got %v", err)
	}
}

func TestGameMakeMoveDetectsWinner(t *testing.T) {
	playerID := uuid.New()
	opponentID := uuid.New()

	game := Game{
		Board: Board{
			{X, X, Empty},
			{O, O, Empty},
			{Empty, Empty, Empty},
		},
		Players: []Player{
			{ID: playerID, Cell: X},
			{ID: opponentID, Cell: O},
		},
		CurrentPlayer: playerID,
		Status:        GameInProgress,
	}

	err := game.MakeMove(playerID, 0, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if game.Status != GameWon {
		t.Fatalf("expected GameWon, got %v", game.Status)
	}

	if game.Winner != playerID {
		t.Fatalf("expected winner %v, got %v", playerID, game.Winner)
	}

	if game.CurrentPlayer != playerID {
		t.Fatalf("current player should not change after game ends")
	}
}

func TestGameJoinPlayer(t *testing.T) {
	player1 := uuid.New()
	player2 := uuid.New()

	game := NewGame()

	if err := game.JoinPlayer(player1, X); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(game.Players) != 1 {
		t.Fatalf("expected 1 player, got %d", len(game.Players))
	}

	if game.Status != GameWaiting {
		t.Fatalf("expected GameWaiting, got %v", game.Status)
	}

	if game.CurrentPlayer != player1 {
		t.Fatalf("expected current player %v, got %v", player1, game.CurrentPlayer)
	}

	if err := game.JoinPlayer(player2, O); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(game.Players) != 2 {
		t.Fatalf("expected 2 players, got %d", len(game.Players))
	}

	if game.Status != GameInProgress {
		t.Fatalf("expected GameInProgress, got %v", game.Status)
	}

	if game.CurrentPlayer != player1 {
		t.Fatalf("expected first player to start, got %v", game.CurrentPlayer)
	}
}

//Тест на повторное присоединение

func TestGameJoinPlayerRejectsDuplicatePlayer(t *testing.T) {
	playerID := uuid.New()

	game := NewGame()

	if err := game.JoinPlayer(playerID, X); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err := game.JoinPlayer(playerID, O)
	if err != ErrPlayerAlreadyJoined {
		t.Fatalf("expected ErrPlayerAlreadyJoined, got %v", err)
	}
}

// Тест на третьего игрока

func TestGameJoinPlayerRejectsThirdPlayer(t *testing.T) {
	player1 := uuid.New()
	player2 := uuid.New()
	player3 := uuid.New()

	game := NewGame()

	if err := game.JoinPlayer(player1, X); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := game.JoinPlayer(player2, O); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err := game.JoinPlayer(player3, X)
	if err != ErrGameFull {
		t.Fatalf("expected ErrGameFull, got %v", err)
	}
}

func TestGameJoinPlayerRejectsDuplicateCell(t *testing.T) {
	player1 := uuid.New()
	player2 := uuid.New()

	game := NewGame()

	if err := game.JoinPlayer(player1, X); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err := game.JoinPlayer(player2, X)
	if err != ErrInvalidPlayerCell {
		t.Fatalf("expected ErrInvalidX, got %v", err)
	}
}
