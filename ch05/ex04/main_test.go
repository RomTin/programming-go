package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestFindLinks(t *testing.T) {

	expected := []string{
		"/",
		"/",
		"#",
		"/doc/",
		"/pkg/",
		"/project/",
		"/help/",
		"/blog/",
		"http://play.golang.org/",
		"#",
		"#",
		"//tour.golang.org/",
		"/dl/",
		"//blog.golang.org/",
		"https://developers.google.com/site-policies#restrictions",
		"/LICENSE",
		"/doc/tos.html",
		"http://www.google.com/intl/en/policies/privacy/",
		"/lib/godoc/style.css",
		"/opensearch.xml",
		"/lib/godoc/jquery.treeview.css",
		"/lib/godoc/jquery.js",
		"/lib/godoc/jquery.treeview.js",
		"/lib/godoc/jquery.treeview.edit.js",
		"/lib/godoc/playground.js",
		"/lib/godoc/godocs.js"}
	file, err := os.Open("./index.html")
	if err != nil {
		t.Errorf("File open error: %v\n", err)
	}
	doc, err := html.Parse(file)
	actual := visit(nil, doc, "a", "href")
	actual = visit(actual, doc, "img", "src")
	actual = visit(actual, doc, "link", "href")
	actual = visit(actual, doc, "script", "src")
	file.Close()

	for i, elem := range expected {
		if elem != actual[i] {
			t.Errorf("something went wrong in processing. %s and %s", elem, actual[i])
		}
	}
}
