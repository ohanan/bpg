package main

import "github.com/ohanan/bpg/pkg/bot"

func main() {
	if err := bot.Run(nil); err != nil {
		panic(err)
	}
}
