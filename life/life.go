package life

import (
	"fmt"
	"math/rand"
)

type Board struct {
	width  int
	height int
	flags  [][]bool
}

func NewBoard(width int, height int) *Board {
	f := make([][]bool, width)
	for x := 0; x < width; x++ {
		f[x] = make([]bool, height)
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
	c += btoi(b.State(x-1, y-1))
	c += btoi(b.State(x, y-1))
	c += btoi(b.State(x+1, y-1))
	c += btoi(b.State(x-1, y))
	c += btoi(b.State(x+1, y))
	c += btoi(b.State(x-1, y+1))
	c += btoi(b.State(x, y+1))
	c += btoi(b.State(x+1, y+1))
	return c
}

func (b *Board) Init(seed int64) {
	r := rand.New(rand.NewSource(seed))
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			b.flags[x][y] = r.Intn(2) == 1
		}
	}
}

func (b *Board) Iterate() *Board {
	i := NewBoard(b.width, b.height)
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			n := b.Neighbours(x, y)
			if n == 2 {
				i.flags[x][y] = b.State(x, y)
			} else if n == 3 {
				i.flags[x][y] = true
			}
		}
	}
	return i
}

func (b *Board) Print() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.State(x, y) {
				fmt.Print("*")
			} else {
				fmt.Print("·")
			}
		}
		fmt.Println()
	}
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
