package main

import (
	"fmt"
	"os"
	"time"

	"appelberg.me/m/life/board"
	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Errorf("could not get term size: %w", err))
	}

	b := board.NewBoard(w, h)
	b.Init(time.Now().UnixNano())

	board.Animate(b, board.AnimationOpts{
		Delay: 100,
	})
}
