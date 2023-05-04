package board

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"time"

	"golang.org/x/term"
)

type AnimationOpts struct {
	Delay time.Duration
}

func resetPosition() {
	fmt.Printf("\033[%d;%dH", 0, 0)
}

func Animate(b *Board, opts AnimationOpts) {
	for {
		resetPosition()
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(fmt.Errorf("could not get term size: %w", err))
		}
		u := b.ToBuffer(w, h)
		_, err = u.WriteTo(os.Stdout)
		if err != nil {
			panic(fmt.Errorf("could not write: %w", err))
		}
		b = b.Iterate()
		time.Sleep(opts.Delay * time.Millisecond)
	}

}

func (b *Board) Print() {
	u := b.ToBuffer(b.width, b.height)
	_, err := u.WriteTo(os.Stdout)
	if err != nil {
		panic(fmt.Errorf("could not print board: %w", err))
	}
}

func (b *Board) ToBuffer(width int, height int) bytes.Buffer {
	var u bytes.Buffer
	w := int(math.Min(float64(width), float64(b.width)))
	h := int(math.Min(float64(height), float64(b.height)))

	for y := 0; y < h; y++ {
		u.WriteString("\033[K") // Clear the line
		if y > 0 {
			u.WriteString("\r\n")
		}
		for x := 0; x < w; x++ {
			if b.State(x, y) {
				u.WriteString("\033[97m*")
			} else {
				u.WriteString("\033[30m·")
			}
		}
	}
	u.WriteString("\033[0m") // Reset color

	return u
}

