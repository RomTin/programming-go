package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestEchoArguments (t *testing.T) {
	buf := &bytes.Buffer{}
	expected := strings.Join(os.Args, " ") + "\n"
	echoArguments(buf)
	actual := buf.String()

	if expected != actual {
		t.Errorf("actual: %v\nexpected: %v", actual, expected)
	}
}