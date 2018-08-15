package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	length := flag.Int("type", 256, "hash length")
	flag.Parse()
	if len(flag.Args()) < 1 {
		return
	}
	fmt.Println(getHash(flag.Args()[0], *length))
}

func getHash(orig string, length int) string {
	switch length {
	case 256:
		return fmt.Sprintf("%x", sha256.Sum256([]byte(orig)))
	case 384:
		return fmt.Sprintf("%x", sha512.Sum384([]byte(orig)))
	case 512:
		return fmt.Sprintf("%x", sha512.Sum512([]byte(orig)))
	default:
		return "invalid hash type"
	}
}
