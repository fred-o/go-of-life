package main

import (
	"fmt"
	"os"
	"time"

	"appelberg.me/m/life/life"
	"golang.org/x/term"
)

func main() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic( fmt.Errorf("could not get term size: %w", err))
	}
	b := life.NewBoard(w, h-1)
	b.Init(time.Now().UnixNano())
	life.Animate(b, life.AnimationOpts{
		Delay: 100,	
	})
}
