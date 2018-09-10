package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {

	expected := map[string]int{
		"body":     1,
		"div":      34,
		"a":        22,
		"input":    1,
		"path":     2,
		"head":     1,
		"script":   10,
		"textarea": 2,
		"pre":      1,
		"select":   1,
		"html":     1,
		"svg":      1,
		"span":     5,
		"button":   1,
		"option":   8,
		"iframe":   1,
		"meta":     3,
		"link":     3,
		"br":       3,
		"title":    2,
		"form":     1}
	file, err := os.Open("./index.html")
	if err != nil {
		t.Errorf("File open error: %v\n", err)
	}
	actual := make(map[string]int)
	doc, err := html.Parse(file)
	outline(actual, doc)
	file.Close()

	for key, val := range expected {
		if val != actual[key] {
			t.Errorf("something went wrong in processing. %d and %d", val, actual[key])
		}
	}
}
