package strategy

import "project03/internal/domain/models"

type MoveStrategy interface {
	NextMove(board models.Board, cell models.Cell) (row, col int, err error)
}
