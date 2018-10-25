// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"../../ch05/links"
)

type Link struct {
	Url   string
	Depth int
}

func crawl(link Link, maxDepth int) []Link {
	if link.Depth >= maxDepth {
		return []Link{}
	}
	fmt.Println(link.Depth, link.Url)
	result := []Link{}
	list, err := links.Extract(link.Url)
	if err != nil {
		log.Print(err)
	}
	for _, url := range list {
		result = append(result, Link{url, link.Depth + 1})
	}
	return result
}

//!+
func main() {
	depth := flag.Int("depth", 10, "depth")
	flag.Parse()
	arglist := []Link{}
	fmt.Println(*depth)
	for _, url := range os.Args[1:] {
		arglist = append(arglist, Link{url, 0})
	}
	worklist := make(chan []Link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- arglist }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, *depth)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[Link]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
