package main

import (
	"testing"
)

func TestEchoArguments(t *testing.T) {
	_, bench1 := echoArgument1()
	_, bench2 := echoArgument2()
	if bench1 == 0 || bench2 == 0 {
		t.Errorf("something went wrong in calculating time")
	}
}