package complex

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func MainC64(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex64(complex(x, y))
			img.Set(px, py, newton64(z))
		}
	}
	png.Encode(out, img)
}

func newton64(z complex64) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(complex128(z*z*z*z-1)) < 1e-6 {
			return color.RGBA{255 - contrast*i, 255, 0, 255}
		}
	}
	return color.Black
}
