package mappers

import (
	dsmodels "project03/internal/datasource/models"
	"project03/internal/domain/models"
)

func ToDomain(dto dsmodels.GameDTO) models.Game {
	var board models.Board

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = models.Cell(dto.Board[i][j])
		}
	}

	players := make([]models.Player, 0, len(dto.Players))

	for _, player := range dto.Players {
		players = append(players, models.Player{
			ID:   player.ID,
			Cell: models.Cell(player.Cell),
		})
	}

	return models.Game{
		ID:            dto.ID,
		Board:         board,
		Players:       players,
		CurrentPlayer: dto.CurrentPlayer,
		Status:        models.GameStatus(dto.Status),
		Winner:        dto.Winner,
		Mode:          models.GameMode(dto.Mode),
	}
}

func ToDTO(game models.Game) dsmodels.GameDTO {
	var board [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = int(game.Board[i][j])
		}
	}

	players := make([]dsmodels.PlayerDTO, 0, len(game.Players))

	for _, player := range game.Players {
		players = append(players, dsmodels.PlayerDTO{
			ID:   player.ID,
			Cell: int(player.Cell),
		})
	}

	return dsmodels.GameDTO{
		ID:            game.ID,
		Board:         board,
		Players:       players,
		CurrentPlayer: game.CurrentPlayer,
		Status:        int(game.Status),
		Winner:        game.Winner,
		Mode:          int(game.Mode),
	}
}
