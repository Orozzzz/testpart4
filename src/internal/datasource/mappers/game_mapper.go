package mappers

import (
    "project03/internal/domain/models"
    dsmodels "project03/internal/datasource/models"
)

func ToDomain(dto dsmodels.GameDTO) models.Game {
    return models.Game{
        ID:    dto.ID,
        Board: dto.Board,
        Turn:  dto.Turn,
    }
}

func ToDTO(game models.Game) dsmodels.GameDTO {
    return dsmodels.GameDTO{
        ID:    game.ID,
        Board: game.Board,
        Turn:  game.Turn,
    }
}