package main

import (
	"flag"
)

const DEFAULT_ANIMAL = "cat"

func main() {
	var animal string
	flag.StringVar(&animal, "animal", DEFAULT_ANIMAL, "animal type")
	flag.StringVar(&animal, "a", DEFAULT_ANIMAL, "animal type")
	flag.Parse()
}
