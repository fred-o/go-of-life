package life

import (
	"fmt"
	"time"
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
		b.Print()
		b = b.Iterate()
		time.Sleep(opts.Delay * time.Millisecond)
	}

}
