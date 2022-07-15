package main

import (
	"context"
	"flag"

	"github.com/ohanan/bpg/internal/common"
	"github.com/ohanan/bpg/internal/panics"
)

func main() {
	var game string
	flag.StringVar(&game, "game", "", "game name")
	var bot string
	flag.StringVar(&bot, "bot", "", "bot name")
	flag.Parse()
	if len(game) == 0 {
		panic("game name is required")
	}
	if len(bot) == 0 {
		panic("bot name is required")
	}
	gc := proto.NewGameClient(panics.If1(common.NewClient(game)))
	bc := proto.NewBotClient(panics.If1(common.NewClient(bot)))
	println(panics.If1(gc.SayHello(context.Background(), &proto.HelloRequest{
		Name: "bpg1",
	})).Message)
	print(panics.If1(bc.SayHello(context.Background(), &proto.HelloRequest{
		Name: "bpg2",
	})).Message)
}
