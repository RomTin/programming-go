package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex04: %v\n", err)
		os.Exit(1)
	}
	links := visit(nil, doc, "a", "href")
	links = visit(links, doc, "img", "src")
	links = visit(links, doc, "link", "href")
	links = visit(links, doc, "script", "src")
	for _, link := range links {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node, tag string, attr string) []string {
	if n.Type == html.ElementNode && n.Data == tag {
		for _, t := range n.Attr {
			if t.Key == attr {
				links = append(links, t.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild, tag, attr)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling, tag, attr)
	}
	return links
}
