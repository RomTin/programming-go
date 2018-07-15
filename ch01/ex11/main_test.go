package main

import (
	"testing"
	"fmt"
)

func TestFetchAll(t *testing.T) {
	urls := [3]string{"https://google.com", "https://yahoo.com", "http://undefined_url.com"}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Sprint(<-ch)
	}
}