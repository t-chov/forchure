package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/t-chov/forchure/lib"
)

const APP_VERSION = "0.0.2"
const DEFAULT_ANIMAL = "cat"

const (
	ExitCodeOK int = iota
	ExitCodeError
	ExitCodeFatal
)

type Cli struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &Cli{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

func (c *Cli) Run(args []string) int {
	var animal string
	var version bool

	flag.BoolVar(&version, "version", false, "show app version")
	flag.StringVar(&animal, "animal", DEFAULT_ANIMAL, "animal type")
	flag.StringVar(&animal, "a", DEFAULT_ANIMAL, "animal type")

	argsBackup := os.Args
	os.Args = args
	flag.Parse()
	os.Args = argsBackup

	if version {
		fmt.Fprintf(c.outStream, "forchure %s\n", APP_VERSION)
		return ExitCodeOK
	}

	buf, err := lib.FetchAnimalTrivia(animal)
	if err != nil {
		fmt.Fprintln(c.errStream, fmt.Errorf("failed to fetch animal trivia: %v", err))
		return ExitCodeError
	}
	trivia, err := lib.ParseAnimalTrivia(buf)
	if err != nil {
		fmt.Fprintln(c.errStream, fmt.Errorf("failed to parse: %v", err))
		return ExitCodeError
	}
	fmt.Fprintln(c.outStream, trivia)
	return ExitCodeError
}
