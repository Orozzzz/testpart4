package service

import (
	"errors"

	"project03/internal/datasource/repository"
	"project03/internal/domain/models"
)

type gameService struct {
	repo repository.GameRepository
}

func NewService(repo repository.GameRepository) GameService {
	return &gameService{
		repo: repo,
	}
}

func (s *gameService) NewGame() models.Game {
	game := models.NewGame()
	s.repo.Save(game)
	return game
}

func (s *gameService) SaveGame(game models.Game) error {
	return s.repo.Save(game)
}

func (s *gameService) GetGame(id string) (models.Game, error) {
	return s.repo.Get(id)
}

func (s *gameService) ValidateBoard(currentGame models.Game, newBoard models.Board) error {
	changes := 0
	var changedRow, changedCol int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if currentGame.Board[i][j] != newBoard[i][j] {
				changes++
				changedRow, changedCol = i, j
			}
		}
	}

	if changes == 0 {
		return errors.New("no moves made")
	}
	if changes > 1 {
		return errors.New("invalid move: more than one cell changed")
	}
	if currentGame.Board[changedRow][changedCol] != 0 {
		return errors.New("cell already occupied")
	}
	if newBoard[changedRow][changedCol] != 1 {
		return errors.New("invalid player symbol")
	}
	return nil
}

func (s *gameService) CheckGameOver(board models.Board) (bool, models.Cell, string) {
	game := models.Game{Board: board}
	return game.IsGameOver()
}

func (s *gameService) GetNextMove(game models.Game) (models.Board, error) {
	if finished, _, _ := game.IsGameOver(); finished {
		return game.Board, errors.New("game is already over")
	}

	newBoard := game.Board
	bestScore := -1000
	bestMove := -1

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if newBoard[i][j] == 0 {
				newBoard[i][j] = 2
				score := s.minimax(newBoard, 0, false)
				newBoard[i][j] = 0
				if score > bestScore {
					bestScore = score
					bestMove = i*3 + j
				}
			}
		}
	}
	if bestMove != -1 {
		row := bestMove / 3
		col := bestMove % 3
		newBoard[row][col] = 2
	}

	return newBoard, nil
}

func (s *gameService) minimax(board models.Board, depth int, isMaximizing bool) int {
	game := models.Game{Board: board}
	if finished, winner, _ := game.IsGameOver(); finished {
		if winner == 2 {
			return 10 - depth
		} else if winner == 1 {
			return -10 + depth
		} else {
			return 0
		}
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = 2
					score := s.minimax(board, depth+1, false)
					board[i][j] = 0
					bestScore = max(bestScore, score)
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == 0 {
					board[i][j] = 1
					score := s.minimax(board, depth+1, true)
					board[i][j] = 0
					bestScore = min(bestScore, score)
				}
			}
		}
		return bestScore
	}
}
