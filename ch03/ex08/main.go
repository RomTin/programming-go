package main

import (
	"os"

	"./complex"
)

func main() {
	complex.MainC128(os.Stdout)
	//complex.MainC64(os.Stdout)
	//complex.MainBigFloat(os.Stdout)
	//complex.MainRat(os.Stdout)
}
