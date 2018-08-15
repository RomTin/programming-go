package main

import "testing"

func TestGetHash1(t *testing.T) {
	expected := "34550715062af006ac4fab288de67ecb44793c3a05c475227241535f6ef7a81b"
	actual := getHash("michael", 256)
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestGetHash2(t *testing.T) {
	expected := "cca73b394d94bdb9b110aaa0a3f4205e0c5f46352d76f7edcd4487428b36e9a58b0776dc5ed957a19bd19d289b02615a"
	actual := getHash("michael", 384)
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}

func TestGetHash3(t *testing.T) {
	expected := "34e1fd6820ce1e79fbbdaae3fc708b634ab1d9765c215b7cd88d4c0c750e87b8c1d478b6112d95ae7bd165f9f73d165263ef7fcee357b48c6bc1f6b591f94ab8"
	actual := getHash("michael", 512)
	if expected != actual {
		t.Errorf("something went wrong in processing.")
	}
}
