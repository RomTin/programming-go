package main

import "fmt"

func duplicate(strings []string) []string {
	i := 0
	buf := ""
	for _, s := range strings {
		if s != buf {
			buf = s
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	data := []string{"one", "one", "one", "two", "three"}
	fmt.Printf("%q\n", duplicate(data))
}
