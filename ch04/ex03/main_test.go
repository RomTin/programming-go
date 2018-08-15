package main

import "testing"

func TestReverse(t *testing.T) {
	origin := [6]int{0, 1, 2, 3, 4, 5}
	copy := &origin
	reverse(&origin)
	if origin[5] != 0 || copy[5] != 0 {
		t.Errorf("something went wrong in processing.")
	}
}
