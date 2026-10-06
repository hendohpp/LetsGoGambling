package main

import (
	"flag"
	"fmt"
)

func main() {
	var genMan bool
	var manPath string

	flag.BoolVar(&genMan, "manual", false, "generate and export GoGambling game manual")
	flag.StringVar(&manPath, "path", "lets-go-gambling-manual.pdf", "output path for game manual pdf")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 && args[0] == "manual" {
		genMan = true
	}

	if genMan {
		fmt.Println("generating GoGambling manual to:", manPath)
	}
}
