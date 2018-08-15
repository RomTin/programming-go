package main

import "fmt"

func rotate(s []int) {
	head := s[0]
	copy(s, s[1:])
	s[len(s)-1] = head
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:])
	fmt.Println(a)
}
