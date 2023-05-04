package board

import (
	"math/rand"
)

type Board struct {
	width  int
	height int
	curr   *[][]bool
	next   *[][]bool
}

func NewBoard(width int, height int) *Board {
	curr := make([][]bool, width)
	for x := 0; x < width; x++ {
		curr[x] = make([]bool, height)
	}
	next := make([][]bool, width)
	for x := 0; x < width; x++ {
		next[x] = make([]bool, height)
	}
	b := Board{width, height, &curr, &next}
	return &b
}

func (b *Board) State(x int, y int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if x >= b.width || y >= b.height {
		return false
	}
	return (*b.curr)[x][y]
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
			(*b.curr)[x][y] = r.Intn(2) == 1
		}
	}
}

func (b *Board) Iterate() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			n := b.Neighbours(x, y)
			if n == 2 {
				(*b.next)[x][y] = b.State(x, y)
			} else if n == 3 {
				(*b.next)[x][y] = true
			} else {
				(*b.next)[x][y] = false
			}
		}
	}
	b.curr, b.next = b.next, b.curr
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
