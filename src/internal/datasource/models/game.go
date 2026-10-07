package models

type GameDTO struct {
    ID    string      `json:"id"`
    Board [3][3]int   `json:"board"`
    Turn  int         `json:"turn"`
}