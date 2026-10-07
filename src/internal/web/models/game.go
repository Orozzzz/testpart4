package models

type GameRequest struct {
	ID    string    `json:"id"`
	Board [3][3]int `json:"board"`
}

type GameResponse struct {
	ID     string    `json:"id"`
	Board  [3][3]int `json:"board"`
	Winner string    `json:"winner,omitempty"`
	Error  string    `json:"error,omitempty"`
}

type NewGameResponse struct {
	ID    string    `json:"id"`
	Board [3][3]int `json:"board"`
	Turn  int       `json:"turn"`
}
