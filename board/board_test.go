package board

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewBoard(t *testing.T) {
	b := NewBoard(48, 32)
	assert.Equal(t, b.width, 48)
	assert.Equal(t, b.height, 32)
}

func TestBoardState(t *testing.T) {
	b := NewBoard(4, 3)
	assert.Equal(t, b.State(0, 0), false)
	assert.Equal(t, b.State(-1, 0), false)
	assert.Equal(t, b.State(0, 10), false)
	(*b.curr)[1][1] = true
	assert.Equal(t, b.State(1, 1), true)
}

func TestBoardNeighbours(t *testing.T) {
	b := NewBoard(4, 3)
	(*b.curr)[1][1] = true
	(*b.curr)[2][1] = true
	assert.Equal(t, b.Neighbours(0, 0), 1)
	assert.Equal(t, b.Neighbours(1, 0), 2)
	assert.Equal(t, b.Neighbours(1, 1), 1)
}

func TestBoardInit(t *testing.T) {
	b := NewBoard(4, 3)
	b.Init(23)
	assert.Equal(t, b.State(0, 0), true)
	assert.Equal(t, b.State(1, 0), false)
}

func TestIterate(t *testing.T) {
	b := NewBoard(4, 3)
	b.Init(23)
	b.Iterate()
	assert.Equal(t, b.State(0, 0), false)
	assert.Equal(t, b.State(1, 0), false)
}

func TestToBuffer(t *testing.T) {
	b := NewBoard(4, 3)
	u := b.ToBuffer(4, 3)
	assert.Equal(t, u.Len(), 101)
	u = b.ToBuffer(2, 2)
	assert.Equal(t, u.Len(), 40)
	u = b.ToBuffer(10, 10)
	assert.Equal(t, u.Len(), 101)
}
