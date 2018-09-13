package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {

	expected1, expected2 := 480, 0
	file, err := os.Open("./index.html")
	if err != nil {
		t.Errorf("File open error: %v\n", err)
	}
	doc, err := html.Parse(file)
	actual1, actual2 := countWordsAndImages(doc)
	file.Close()

	if expected1 != actual1 || expected2 != actual2 {
		t.Errorf("something went wrong in processing.")
	}
}
