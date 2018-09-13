package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	htmls := []string{
		"<html><body><img src=\"sample.img\"/></body></html>",
		"<html><body><a href=\"sample.link\">hoge-fuga</a></body></html>",
		"<html><body><p class=\"sample.class\">sentence</p></body></html>",
		"<html><body><script>var a = 1; alert(a);</script></body></html>",
	}
	for _, ht := range htmls {
		reader := strings.NewReader(ht)
		h, err := html.Parse(reader)
		forEachNode(html.Parse(h), startElement, endElement)
	}

}
