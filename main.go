package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fred-o/go-of-life/board"
	"github.com/fred-o/go-of-life/ui"
	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(fmt.Errorf("could not get term size: %w", err))
	}

	b := board.NewBoard(w, h)
	b.Init(time.Now().UnixNano())

	ui.Animate(b, ui.AnimationOpts{
		Delay: 150,
	})
}
