package mappers

import (
    "project03/internal/domain/models"
    webmodels "project03/internal/web/models"
)

func ToWebResponse(game models.Game, winner string) webmodels.GameResponse {
    return webmodels.GameResponse{
        ID:     game.ID,
        Board:  game.Board,
        Winner: winner,
    }
}

func ToDomainRequest(req webmodels.GameRequest) models.Game {
    return models.Game{
        ID:    req.ID,
        Board: req.Board,
        Turn:  1, 
    }
}

func ToNewGameResponse(game models.Game) webmodels.NewGameResponse {
    return webmodels.NewGameResponse{
        ID:    game.ID,
        Board: game.Board,
        Turn:  game.Turn,
    }
}