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
	b := NewBoard(3, 3)
	assert.Equal(t, b.State(0, 0), false)
	assert.Equal(t, b.State(-1, 0), false)
	assert.Equal(t, b.State(0, 10), false)
	b.flags[1][1] = true
	assert.Equal(t, b.State(1, 1), true)
}

func TestBoardNeighbours(t *testing.T) {
	b := NewBoard(3, 3)
	b.flags[1][1] = true
	b.flags[2][1] = true
	assert.Equal(t, b.Neighbours(0, 0), 1)
	assert.Equal(t, b.Neighbours(1, 0), 2)
	assert.Equal(t, b.Neighbours(1, 1), 1)
}

func TestBoardIni(t *testing.T) {
	b := NewBoard(3, 3)
	b.Init(23)
	assert.Equal(t, b.State(0, 0), true)
	assert.Equal(t, b.State(1, 0), false)
}
