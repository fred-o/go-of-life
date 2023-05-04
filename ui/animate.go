package ui

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/fred-o/go-of-life/board"
	"golang.org/x/term"
)

type AnimationOpts struct {
	Delay time.Duration
}

func resetPosition() {
	fmt.Printf("\033[%d;%dH", 0, 0)
}

func Animate(b *board.Board, opts AnimationOpts) {
	for {
		resetPosition()
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(fmt.Errorf("could not get term size: %w", err))
		}
		r := Render(b, w, h)
		r.Draw()
		b.Iterate()
		time.Sleep(opts.Delay * time.Millisecond)
	}

}

type Rendered struct {
	buf bytes.Buffer
}

func Render(b *board.Board, width int, height int) *Rendered {
	var u bytes.Buffer
	w := int(math.Min(float64(width), float64(b.Width)))
	h := int(math.Min(float64(height), float64(b.Height)))

	for y := 0; y < h; y++ {
		u.WriteString("\033[K") // Clear the line
		if y > 0 {
			u.WriteString("\r\n")
		}
		for x := 0; x < w; x++ {
			if b.State(x, y) {
				u.WriteString("\033[97m×")
			} else {
				u.WriteString("\033[30m·")
			}
		}
	}
	u.WriteString("\033[0m") // Reset color

	return &Rendered{u}
}

func (r *Rendered) Draw() {
	_, err := r.buf.WriteTo(os.Stdout)
	if err != nil {
		panic(fmt.Errorf("could not write: %w", err))
	}
}
