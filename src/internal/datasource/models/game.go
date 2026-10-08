package models

import "github.com/google/uuid"

type PlayerDTO struct {
	ID   uuid.UUID `json:"id"`
	Cell int       `json:"cell"`
}

type GameDTO struct {
	ID            string      `json:"id"`
	Board         [3][3]int   `json:"board"`
	Players       []PlayerDTO `json:"players"`
	CurrentPlayer uuid.UUID   `json:"current_player"`
	Status        int         `json:"status"`
	Winner        uuid.UUID   `json:"winner"`
	Mode          int         `json:"mode"`
}
