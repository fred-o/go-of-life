package life

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
	b.flags[1][1] = true
	assert.Equal(t, b.State(1, 1), true)
}

func TestBoardNeighbours(t *testing.T) {
	b := NewBoard(4, 3)
	b.flags[1][1] = true
	b.flags[2][1] = true
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
	b1 := NewBoard(4, 3)
	b1.Init(23)
	b2 := b1.Iterate()
	assert.Equal(t, b2.State(0, 0), false)
	assert.Equal(t, b2.State(1, 0), false)
}

func TestToBuffer(t *testing.T) {
	b := NewBoard(4, 3)
	u := b.ToBuffer()
	assert.Equal(t, u.Len(), 28)
}
