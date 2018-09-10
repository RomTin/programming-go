package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	texts(doc)
}

func texts(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(strings.Replace(n.Data, "\n", "", -1))
	} else if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			texts(c)
		}
	}
}
