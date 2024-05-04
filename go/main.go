package main

import (
	"flag"
	"log"
)

func main() {
	inputFile := flag.String("file", "", "Measurements input file")
	impl := flag.String("impl", "brute_force", "Implementation to use [brute_force]")

	flag.Parse()

	if *inputFile == "" {
		log.Panic("Measurements input file requred!")
	}

	switch *impl {
	case "brute_force":
		BruteForceSolution(*inputFile)
	default:
		log.Panic("Invalid input for impl!")
	}
}
