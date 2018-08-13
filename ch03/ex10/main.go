package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s) % 3
	if n == 0 {
		n = 3
	}
	fmt.Fprintf(&buf, "%s", s[:n])
	for ; n < len(s); n += 3 {
		fmt.Fprintf(&buf, ",%s", s[n:(n+3)])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1000000"))
}
