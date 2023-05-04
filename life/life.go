package life

import "fmt"


type Board struct {
	width int
	height int
	flags [][]bool
}

func NewBoard(width int, height int) *Board {
	f := make([][]bool, height)
	fmt.Println(f)
	for i := 0; i < height; i++ {
		f[i] = make([]bool, width)
	}
	b := Board{width, height, f}
	return &b
}

func (b *Board) State(x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= b.width || y >= b.height {
		return false
	}
	return b.flags[x][y]
}

func (b *Board) Iterate() *Board {
	return b
}

