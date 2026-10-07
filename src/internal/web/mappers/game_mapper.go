package mappers

import (
    "project03/internal/domain/models"
    webmodels "project03/internal/web/models"
)

func ToWebResponse(game models.Game, winner string) webmodels.GameResponse {
	var board [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = int(game.Board[i][j])
		}
	}

	return webmodels.GameResponse{
		ID:     game.ID,
		Board:  board,
		Winner: winner,
	}
}


func ToDomainRequest(req webmodels.GameRequest) models.Game {
	var board models.Board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = models.Cell(req.Board[i][j])
		}
	}

	return models.Game{
		ID:    req.ID,
		Board: board,
		Turn:  1,
	}
}


func ToNewGameResponse(game models.Game) webmodels.NewGameResponse {
	var board [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = int(game.Board[i][j])
		}
	}

	return webmodels.NewGameResponse{
		ID:    game.ID,
		Board: board,
		Turn:  game.Turn,
	}
}


func ToDomainBoard(board [3][3]int) models.Board {
	var result models.Board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = models.Cell(board[i][j])
		}
	}

	return result
}