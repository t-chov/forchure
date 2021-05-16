package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunWhenVersionOption(t *testing.T) {
	out := new(bytes.Buffer)
	cli := &Cli{outStream: out, errStream: new(bytes.Buffer)}
	args := strings.Split("forchure -version", " ")

	if status := cli.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d. but found %d", ExitCodeOK, status)
	}

	expectedOutput := fmt.Sprintf("forchure %s\n", APP_VERSION)
	actualOutput := out.String()
	if actualOutput != expectedOutput {
		t.Errorf("expected %s. but found %s", expectedOutput, actualOutput)
	}
}
