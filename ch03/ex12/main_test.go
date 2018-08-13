package main

import (
	"testing"
)

func TestAnagram1(t *testing.T) {
	expected := true
	actual := isAnagram("abc", "bca")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestAnagram2(t *testing.T) {
	expected := true
	actual := isAnagram("aabbcc", "cbacba")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestAnagram3(t *testing.T) {
	expected := true
	actual := isAnagram("墾田永年私財法", "田財私永法年墾")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestAnagram4(t *testing.T) {
	expected := false
	actual := isAnagram("aaa", "bbb")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
