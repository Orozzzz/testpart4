package models

type Board [3][3]int

func NewBoard() Board {
	return Board{}
}

func (b Board) IsEmpty(row, col int) bool {
	return b[row][col] == 0
}

func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) MakeMove(row, col int, player int) bool {
	if b[row][col] != 0 {
		return false
	}
	b[row][col] = player
	return true
}
