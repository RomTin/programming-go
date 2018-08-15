package main

import "fmt"

func rotate(s []int, n int) {
	switch n {
	case 0:
		break
	default:
		head := s[0]
		copy(s, s[1:])
		s[len(s)-1] = head
		rotate(s, n-1)
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:], 3)
	fmt.Println(a)
}
