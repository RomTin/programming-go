package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

// e.g.) http://localhost:8000/?mag=1024&px=2&py=2
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mag, _ := strconv.ParseInt(r.FormValue("mag"), 10, 64)
	px, _ := strconv.ParseFloat(r.FormValue("px"), 64)
	py, _ := strconv.ParseFloat(r.FormValue("py"), 64)
	draw(w, mag, px, py)
}

func draw(out io.Writer, mag int64, px float64, py float64) {
	xmin, ymin, xmax, ymax := -px, -py, +px, +py
	width, height := int(mag), int(mag)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := 0; py < int(height); py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(out, img)
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.RGBA{255 - contrast*i, 255, 0, 255}
		}
	}
	return color.Black
}
