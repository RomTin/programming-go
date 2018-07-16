package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"math/rand"
	"time"
	"io/ioutil"
)

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
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