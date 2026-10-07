package models

type Cell int

const (
	Empty Cell = iota
	PlayerCell
	ComputerCell
)


type Board [3][3]Cell

func NewBoard() Board {
	return Board{}
}

func (b Board) IsEmpty(row, col int) bool {
	return b[row][col] == Empty
}

func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func (b *Board) MakeMove(row, col int, cell Cell) bool {
	if b[row][col] != Empty {
		return false
	}
	b[row][col] = cell
	return true
}
