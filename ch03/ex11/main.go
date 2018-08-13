package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	p := ""
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		buf.WriteString(s[:1])
		s = s[1:]
	}
	dot := strings.Index(s, ".")
	if dot != -1 {
		s, p = s[:dot], s[dot:]
	}
	n := len(s) % 3
	if n == 0 {
		n = 3
	}
	fmt.Fprintf(&buf, "%s", s[:n])
	for ; n < len(s); n += 3 {
		fmt.Fprintf(&buf, ",%s", s[n:(n+3)])
	}
	buf.WriteString(p)
	return buf.String()
}

func main() {
	fmt.Println(comma("-1000000.12345"))
}
