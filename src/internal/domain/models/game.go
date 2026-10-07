package models

import "github.com/google/uuid"

type Game struct {
	ID    string
	Board Board
	Turn  int
}

func NewGame() Game {
	return Game{
		ID:    uuid.New().String(),
		Board: NewBoard(),
		Turn:  1,
	}
}

func (g Game) CheckWinner() (int, bool) {
	board := g.Board

	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0], true
		}
	}

	for j := 0; j < 3; j++ {
		if board[0][j] != 0 && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return board[0][j], true
		}
	}

	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0], true
	}
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2], true
	}

	return 0, false
}

func (g Game) IsGameOver() (bool, int, string) {
	if winner, ok := g.CheckWinner(); ok {
		result := "draw"
		if winner == 1 {
			result = "player"
		} else {
			result = "computer"
		}
		return true, winner, result
	}

	if g.Board.IsFull() {
		return true, 0, "draw"
	}
	return false, 0, ""
}
