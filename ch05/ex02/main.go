package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	dic := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(dic, doc)
	for key, val := range dic {
		fmt.Printf("%-10s :%3d\n", key, val)
	}
}

func outline(count map[string]int, n *html.Node) {
	if n.Type == html.TextNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(count, c)
	}
}
