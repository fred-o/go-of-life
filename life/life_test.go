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
}
