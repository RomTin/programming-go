package main

import "testing"

func TestWorkfreq(t *testing.T) {
	expected := map[string]int{"dolore": 2, "ea": 1, "adipiscing": 1, "incididunt": 1, "ut": 2, "aliqua.": 1, "aute": 1, "sunt": 1, "sit": 1, "labore": 1, "et": 1, "cupidatat": 1, "ipsum": 1, "dolor": 2, "amet,": 1, "culpa": 1, "deserunt": 1, "anim": 1, "laborum.": 1, "do": 1, "exercitation": 1, "nisi": 1, "esse": 1, "Lorem": 1, "minim": 1, "irure": 1, "commodo": 1, "reprehenderit": 1, "laboris": 1, "sint": 1, "cillum": 1, "eu": 1, "voluptate": 1, "velit": 1, "eiusmod": 1, "Duis": 1, "in": 3, "tempor": 1, "Excepteur": 1, "proident,": 1, "occaecat": 1, "quis": 1, "nostrud": 1, "aliquip": 1, "ullamco": 1, "consequat.": 1, "nulla": 1, "officia": 1, "id": 1, "sed": 1, "Ut": 1, "enim": 1, "qui": 1, "fugiat": 1, "non": 1, "consectetur": 1, "ad": 1, "ex": 1, "mollit": 1, "elit,": 1, "magna": 1, "pariatur.": 1, "veniam,": 1, "est": 1}
	actual := wordfreq("./sample.txt")
	for k, xv := range actual {
		if yv, ok := expected[k]; !ok || yv != xv {
			t.Errorf("actual count of %s is not correct.", k)
		}
	}
}