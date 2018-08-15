package main

import (
	"testing"
)

func TestDuplicateSpace1(t *testing.T) {
	origin := []byte("hoge \tfuga \tpiyo")
	fixed := duplicateSpace(origin)
	if string(fixed) != "hoge fuga piyo" {
		t.Errorf("something went wrong in processing.")
	}
}
