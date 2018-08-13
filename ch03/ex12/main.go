package main

import (
	"fmt"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool { return s[i] < s[j] }
func (s sortRunes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortRunes) Len() int           { return len(s) }
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func isAnagram(a, b string) bool {
	return SortString(a) == SortString(b)
}

func main() {
	fmt.Println(isAnagram("abc", "bca"))
	fmt.Println(isAnagram("aaa", "aaa"))
	fmt.Println(isAnagram("aabbcc", "cbacba"))
	fmt.Println(isAnagram("123", "321"))
	fmt.Println(isAnagram("RomTin", "niTmoR"))
	fmt.Println(isAnagram("墾田永年私財法", "田財私永法年墾"))
}
