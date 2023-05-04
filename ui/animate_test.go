package ui

import (
	"testing"

	"github.com/fred-o/go-of-life/board"
	"gotest.tools/assert"
)

func TestToBuffer(t *testing.T) {
	b := board.NewBoard(4, 3)
	r := Render(b, 4, 3)
	assert.Equal(t, r.buf.Len(), 101)
	r = Render(b, 2, 2)
	assert.Equal(t, r.buf.Len(), 40)
	r = Render(b, 10, 10)
	assert.Equal(t, r.buf.Len(), 101)
}
