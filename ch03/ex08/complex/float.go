package complex

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
)

type ComplexF struct {
	r *big.Float
	i *big.Float
}

func plusF(a *ComplexF, b *ComplexF) *ComplexF {
	return &ComplexF{
		new(big.Float).Add(a.r, b.r),
		new(big.Float).Add(a.i, b.i)}
}

func minusF(a *ComplexF, b *ComplexF) *ComplexF {
	return &ComplexF{
		new(big.Float).Sub(a.r, b.r),
		new(big.Float).Sub(a.i, b.i)}
}

func multiF(a *ComplexF, b *ComplexF) *ComplexF {
	return &ComplexF{
		new(big.Float).Sub(
			new(big.Float).Mul(a.r, b.r),
			new(big.Float).Mul(a.i, b.i)),
		new(big.Float).Add(
			new(big.Float).Mul(a.r, b.i),
			new(big.Float).Mul(a.i, b.r))}
}

func divF(a *ComplexF, b *ComplexF) *ComplexF {
	bunbo := new(big.Float).Add(
		new(big.Float).Mul(b.r, b.r),
		new(big.Float).Mul(b.i, b.i))
	new_r := new(big.Float).Quo(
		new(big.Float).Add(
			new(big.Float).Mul(a.r, b.r),
			new(big.Float).Mul(a.i, b.i)),
		bunbo)
	new_i := new(big.Float).Quo(
		new(big.Float).Sub(
			new(big.Float).Mul(b.r, a.i),
			new(big.Float).Mul(a.r, b.i)),
		bunbo)
	return &ComplexF{new_r, new_i}
}

func pdistF(a *ComplexF) *big.Float {
	return new(big.Float).Add(
		new(big.Float).Mul(a.r, a.r),
		new(big.Float).Mul(a.i, a.i))
}

func MainBigFloat(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, newtonF(&ComplexF{
				new(big.Float).SetFloat64(x),
				new(big.Float).SetFloat64(y)}))
		}
	}
	png.Encode(out, img)
}

func newtonF(z *ComplexF) color.Color {
	const iterations = 37
	const contrast = 7
	if pdistF(z).Cmp(new(big.Float).SetFloat64(0)) == -1 {
		return color.Black
	}
	for i := uint8(0); i < iterations; i++ {
		z = minusF(z, divF(
			divF(
				&ComplexF{
					new(big.Float).SetFloat64(1),
					new(big.Float).SetFloat64(0)},
				multiF(z, multiF(z, z))),
			&ComplexF{
				new(big.Float).SetFloat64(4),
				new(big.Float).SetFloat64(0)}))
		if pdistF(minusF(multiF(z, multiF(z, multiF(z, z))), &ComplexF{
			new(big.Float).SetFloat64(1),
			new(big.Float).SetFloat64(0)})).Cmp(
			new(big.Float).SetFloat64(1e-12)) == -1 {
			return color.RGBA{255 - contrast*i, 255, 0, 255}
		}
	}
	return color.Black
}
