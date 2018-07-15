package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)

func main() {
  for _, url := range os.Args[1:] {
    fetch(url)
  }
}

func fetch(url string) string {
  if !strings.HasPrefix(url, "http://") {
    url = "http://" + url
  }
  resp, err := http.Get(url)
  if err != nil {
    fmt.Fprintf(os.Stderr, "error in fetch: %v\n", err)
    os.Exit(1)
  }
  _, err = io.Copy(os.Stdout, resp.Body)
  status := resp.Status
  fmt.Printf("HTTP Status Code: %v\n", resp.Status)
  resp.Body.Close()
  if err != nil {
    fmt.Fprintf(os.Stderr, "error in fetch: reading %s: %v\n", url, err)
    os.Exit(1)
  }
  return status
}