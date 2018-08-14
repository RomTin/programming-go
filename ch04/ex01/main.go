package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte("sample1"))
	c2 := sha256.Sum256([]byte("sample2"))
	if len(os.Args) > 2 {
		c1 = sha256.Sum256([]byte(os.Args[1]))
		c2 = sha256.Sum256([]byte(os.Args[2]))
	}
	res := 0

	for i := 0; i < 32; i++ {
		res += countBits(c1[i], c2[i])
	}
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(res)
}

func countBits(a uint8, b uint8) int {
	res := 0
	for i := 0; i < 8; i++ {
		if (a & uint8(1)) == (b & uint8(1)) {
			res++
		}
		a >>= 1
		b >>= 1
	}
	return res
}
