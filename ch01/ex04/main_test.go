package main

import (
	"fmt"
	"os"
	"testing"
)

func TestCountLines(t *testing.T) {
	counts := make(map[string]map[string]int)
	files := [3]string{"./text/sample1.txt", "./text/sample2.txt", "./text/sample3.txt"}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex04_test: %v\n", err)
			continue
	   	}
	   	countLines(f, counts)
	   	f.Close()
	}

	if !check1and2(counts) || !check1and3(counts) || !check2and3(counts) || !check1and2and3(counts) {
		t.Errorf("Something went wrong in counting lines or input text files.")
	}
}

func check1and2(counts map[string]map[string]int) bool {
	line := "This line is contained in 1 and 2."
	return counts[line]["./text/sample1.txt"] == 1 && counts[line]["./text/sample2.txt"] == 1
}

func check1and3(counts map[string]map[string]int) bool {
	line := "This line is contained in 1 and 3."
	return counts[line]["./text/sample1.txt"] == 1 && counts[line]["./text/sample3.txt"] == 1
}

func check2and3(counts map[string]map[string]int) bool {
	line := "This line is contained in 2 and 3."
	return counts[line]["./text/sample2.txt"] == 1 && counts[line]["./text/sample3.txt"] == 1
}

func check1and2and3(counts map[string]map[string]int) bool {
	line := "This line is contained in 1, 2 and 3."
	return counts[line]["./text/sample1.txt"] == 1 && counts[line]["./text/sample2.txt"] == 1 && counts[line]["./text/sample3.txt"] == 1
}