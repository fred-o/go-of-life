package life

type Board struct {
	width  int
	height int
	flags  [][]bool
}

func NewBoard(width int, height int) *Board {
	f := make([][]bool, height)
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

func (b *Board) Neighbours(x int, y int) int {
	c := 0
	c += btoi(b.State(x - 1, y - 1))
	c += btoi(b.State(x, y - 1))
	c += btoi(b.State(x + 1, y - 1))

	c += btoi(b.State(x - 1, y))
	c += btoi(b.State(x + 1, y))

	c += btoi(b.State(x - 1, y + 1))
	c += btoi(b.State(x, y + 1))
	c += btoi(b.State(x + 1, y + 1))
	return c
}

func (b *Board) Iterate() *Board {
	return b
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
