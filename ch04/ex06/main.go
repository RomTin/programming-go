package main

import (
	"fmt"
	"unicode"
)

func duplicateSpace(line []byte) []byte {
	i := 0
	spaceFlag := false
	for _, s := range line {
		if !spaceFlag && unicode.IsSpace(rune(s)) {
			line[i] = ' '
			spaceFlag = true
			i++
		} else if !unicode.IsSpace(rune(s)) {
			line[i] = s
			spaceFlag = false
			i++
		}
	}
	return line[:i]
}

func main() {
	fmt.Printf("%q\n", duplicateSpace([]byte("hoge \tfuga \npiyo")))
}
