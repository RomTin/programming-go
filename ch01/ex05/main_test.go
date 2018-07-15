package main

import (
	"testing"
	"math/rand"
	"bytes"
	"time"
	"image/gif"
)

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	gif_buffer := new(bytes.Buffer)
	lissajous(gif_buffer)
	r := bytes.NewReader(gif_buffer.Bytes())
	_, err := gif.Decode(r)
	if err != nil {
		t.Errorf("something went wrong in generating GIF")
	}
}
