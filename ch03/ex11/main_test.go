package main

import (
	"testing"
)

func TestComma1(t *testing.T) {
	expected := "-1,000,000"
	actual := comma("-1000000")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
func TestComma2(t *testing.T) {
	expected := "+1,000,000"
	actual := comma("+1000000")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
func TestComma3(t *testing.T) {
	expected := "100.123"
	actual := comma("100.123")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
func TestComma4(t *testing.T) {
	expected := "-1,000.12345"
	actual := comma("-1000.12345")
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
