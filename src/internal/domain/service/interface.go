package service

import "project03/internal/domain/models"

type GameService interface {
    GetNextMove(game models.Game) (models.Board, error)
    ValidateBoard(currentGame models.Game, newBoard models.Board) error
    CheckGameOver(board models.Board) (bool, models.Cell, string)
    NewGame() models.Game
    SaveGame(game models.Game) error
    GetGame(id string) (models.Game, error)
}