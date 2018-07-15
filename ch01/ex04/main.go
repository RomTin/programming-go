package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]map[string]int)
  files := os.Args[1:]
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range files {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }
  for line, files := range counts {
    fmt.Println(line)
    for filename, n := range files {
      fmt.Printf("\t%d\t%s\n", n, filename)
    }
  }
}

func countLines(f *os.File, counts map[string]map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    if counts[input.Text()] == nil {
      counts[input.Text()] = make(map[string]int)
    }
    counts[input.Text()][f.Name()]++
  }
}
