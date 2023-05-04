package main

import (
	"bufio"
	"os"

	"appelberg.me/m/life/life"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	b := life.NewBoard(48, 32)
	b.Init(23)
	for ;; {
		b.Print()
		b = b.Iterate()
		_, _, _ = reader.ReadRune()
	}
}
