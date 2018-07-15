package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	expected := "200 OK"
	actual := fetch("RomTin.github.io")
	if expected != actual {
		t.Errorf("something went wrong in http request")
	}
}