package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handler))

	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Errorf("something went wrong in request. %v", err)
	}

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("something went wrong in reading data. %v", err)
	}
}
