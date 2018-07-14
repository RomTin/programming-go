package main

import (
  "fmt"
  "io"
  "os"
  "strings"
)

func echoArguments(w io.Writer) {
	fmt.Fprint(w, strings.Join(os.Args, " ") + "\n")
}

func main() {
	echoArguments(os.Stdout)
}
