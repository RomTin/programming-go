package main

import "testing"

func TestDuplicate(t *testing.T) {
	origin := []string{"one", "one", "two", "three", "two"}
	fixed := duplicate(origin)
	if len(origin) != 5 || len(fixed) != 4 {
		t.Errorf("something went wrong in processing.")
	}
}
