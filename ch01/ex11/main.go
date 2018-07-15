package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
  "strings"
)

func main() {
  start := time.Now()
  file, err := os.Create(fmt.Sprintf("./%v.log", start))
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v", err)
  }
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    if !strings.HasPrefix(url, "http://") {
      url = "http://" + url
    }  
    go fetch(url, ch)
  }
  for range os.Args[1:] {
    file.Write(([]byte)(fmt.Sprintf("%s\n", <-ch)))
  }
  file.Write(([]byte)(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())))
  file.Close()
}

func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err)
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
