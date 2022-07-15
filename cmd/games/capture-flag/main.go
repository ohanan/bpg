package main

import (
	"github.com/ohanan/bpg/pkg/game"
)

func main() {
	if err := game.Run(nil); err != nil {
		panic(err)
	}
}
