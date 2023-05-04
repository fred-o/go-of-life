package life

import (
	"fmt"
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
