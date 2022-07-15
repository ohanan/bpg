package main

import (
	"flag"

	"github.com/ohanan/bpg/internal/workspace"
)

func parseFlags() {
	workspace.RegisterFlags()
	flag.Parse()
}

func main() {
	parseFlags()
}
