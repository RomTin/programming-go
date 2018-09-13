package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/adonovan/gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
	a, _ := url.Parse(worklist[0])
	allowed := a.Host
	for _, work := range worklist {
		u, _ := url.Parse(work)
		if u.Host == allowed {
			if err := os.MkdirAll("./"+u.Path, 0777); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)

	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
