package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	tmp_w, tmp_i := 0, 0
	if n.Type == html.TextNode {
		words += count(n.Data)
	} else if n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tmp_w, tmp_i = countWordsAndImages(c)
		words += tmp_w
		images += tmp_i
	}
	return
}

func count(i string) (words int) {
	scanner := bufio.NewScanner(strings.NewReader(i))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	return
}

func main() {
	words, images, err := CountWordsAndImages("https://www.pixiv.net/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex04: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(words, images)
}
