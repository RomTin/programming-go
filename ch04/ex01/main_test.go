package main

import "testing"

func TestCountBits(t *testing.T) {
	expected := 8
	actual := countBits(uint8(128), uint8(128))
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestMain(t *testing.T) {
	main()
}
