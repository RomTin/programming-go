package main

import "testing"

func TestReverse(t *testing.T) {
	expected := "gfedcba"
	origin := []byte("abcdefg")
	reverse(origin)
	if string(origin) != expected {
		t.Errorf("something went wrong in processing.")
	}
}
