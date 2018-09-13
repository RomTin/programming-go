package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	sample := "abc$foodef$foo"
	funcs := []func(string) string{
		func(a string) string { return a },
		func(a string) string { return "Foo" },
		func(a string) string { return strings.ToUpper(a) },
		func(a string) string { return strings.ToLower(a) },
	}
	expected := []string{
		"abcfoodeffoo",
		"abcFoodefFoo",
		"abcFOOdefFOO",
		"abcfoodeffoo",
	}
	for index, f := range funcs {
		if expected[index] != expand(sample, f) {
			t.Errorf("something went wrong in processing. %s and %s", expected[index], expand(sample, f))
		}
	}
}
