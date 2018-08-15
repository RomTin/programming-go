package main

import "testing"

func TestRotate(t *testing.T) {
	origin := [4]int{0, 1, 2, 3}
	rotate(origin[:])
	if origin[3] != 0 {
		t.Errorf("something went wrong in processing.")
	}
	rotate(origin[:])
	if origin[3] != 1 {
		t.Errorf("something went wrong in processing.")
	}
}
