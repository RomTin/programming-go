package main

import "testing"

func TestRotate(t *testing.T) {
	origin := [4]int{0, 1, 2, 3}
	rotate(origin[:], 2)
	if origin[3] != 1 {
		t.Errorf("something went wrong in processing.")
	}
	rotate(origin[:], 2)
	if origin[3] != 3 {
		t.Errorf("something went wrong in processing.")
	}
}
