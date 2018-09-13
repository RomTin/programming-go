package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(expand("abc$foodef$foo", func(s string) string { return "FOO" }))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
