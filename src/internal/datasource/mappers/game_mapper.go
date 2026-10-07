package mappers

import (
    "project03/internal/domain/models"
    dsmodels "project03/internal/datasource/models"
)

func ToDomain(dto dsmodels.GameDTO) models.Game {
    var board models.Board

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = models.Cell(dto.Board[i][j])
        }
    }

    return  models.Game{
        ID: dto.ID,
        Board: board,
        Turn: dto.Turn,
    }
}

func ToDTO(game models.Game) dsmodels.GameDTO {
    var board [3][3]int 

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[i][j] = int(game.Board[i][j])
        }
    }

    return  dsmodels.GameDTO{
        ID: game.ID,
        Board: board,
        Turn: game.Turn,
    }
}