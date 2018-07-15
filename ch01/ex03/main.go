package main

import (
  "fmt"
  "os"
  "strings"
  "time"
)

func echoArgument1() (string, time.Duration) {
  begin := time.Now()
  var s, sep string
  for i := 1; i < len(os.Args[1:]); i++ {
    s += sep + os.Args[i]
    sep = " "
  }
  end := time.Now()
  return s, end.Sub(begin)
}

func echoArgument2() (string, time.Duration) {
  begin := time.Now()
  s := strings.Join(os.Args[1:], " ")
  end := time.Now()
  return s, end.Sub(begin)
}

func main() {
  _, bench1 := echoArgument1()
  _, bench2 := echoArgument2()
  fmt.Printf("引数の数: %d\n", len(os.Args[1:]))
  fmt.Printf("効率悪いほう: %v\n", bench1)
  fmt.Printf("効率良いほう: %v\n", bench2)
}
