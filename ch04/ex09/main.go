package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	dict := wordfreq(os.Args[1])
	fmt.Printf("\tword\tcount\n")
	for key, count := range dict {
		fmt.Printf("%-15s\t%d\n", key, count)
	}
}

func wordfreq(path string) map[string]int {
	dict := make(map[string]int)
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return dict
	}
	f := bufio.NewScanner(file)
	f.Split(bufio.ScanWords)
	for f.Scan() {
		word := f.Text()
		dict[word]++
	}
	file.Close()
	return dict
}
