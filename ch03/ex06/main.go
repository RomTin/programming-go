package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	dif := 1.0 / 500
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			za := mandelbrot(complex(x, y))
			zb := mandelbrot(complex(x+dif, y))
			zc := mandelbrot(complex(x, y+dif))
			zd := mandelbrot(complex(x+dif, y+dif))
			img.Set(px, py, ss([]color.Color{za, zb, zc, zd}))
		}
	}
	png.Encode(os.Stdout, img)
}

func ss(px []color.Color) color.Color {
	var Red, Green, Blue, Alpha uint32 = 0, 0, 0, 0
	for _, p := range px {
		r, g, b, a := p.RGBA()
		Red += r
		Green += g
		Blue += b
		Alpha += a
	}
	Red /= 4
	Green /= 4
	Blue /= 4
	Alpha /= 4
	return color.RGBA{uint8(Red), uint8(Green), uint8(Blue), uint8(Alpha)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - n*30, 0, 0, 255}
		}
	}
	return color.Black
}
