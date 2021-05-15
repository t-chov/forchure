package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/t-chov/forchure/lib"
)

const DEFAULT_ANIMAL = "cat"

func main() {
	var animal string
	flag.StringVar(&animal, "animal", DEFAULT_ANIMAL, "animal type")
	flag.StringVar(&animal, "a", DEFAULT_ANIMAL, "animal type")
	flag.Parse()

	buf, err := lib.FetchAnimalTrivia(animal)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to fetch animal trivia: %v", err))
		os.Exit(1)
	}
	trivia, err := lib.ParseAnimalTrivia(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to parse: %v", err))
		os.Exit(1)
	}
	fmt.Println(trivia)
}
