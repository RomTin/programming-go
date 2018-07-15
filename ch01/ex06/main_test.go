package main

import (
	"testing"
	"math/rand"
	"bytes"
	"time"
	"image/gif"
	"image/color"
)

func TestLissajous(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		palette = append(palette, color.RGBA{uint8(rand.Int63() % 255), uint8(rand.Int63() % 255), uint8(rand.Int63() % 255), 0xFF})
	}	
	gif_buffer := new(bytes.Buffer)
	lissajous(gif_buffer)
	r := bytes.NewReader(gif_buffer.Bytes())
	_, err := gif.Decode(r)
	if err != nil {
		t.Errorf("something went wrong in generating GIF")
	}
}
