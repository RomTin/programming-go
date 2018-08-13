package complex

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
)

type ComplexR struct {
	r *big.Rat
	i *big.Rat
}

func plus(a *ComplexR, b *ComplexR) *ComplexR {
	return &ComplexR{
		new(big.Rat).Add(a.r, b.r),
		new(big.Rat).Add(a.i, b.i)}
}

func minus(a *ComplexR, b *ComplexR) *ComplexR {
	return &ComplexR{
		new(big.Rat).Sub(a.r, b.r),
		new(big.Rat).Sub(a.i, b.i)}
}

func multi(a *ComplexR, b *ComplexR) *ComplexR {
	return &ComplexR{
		new(big.Rat).Sub(
			new(big.Rat).Mul(a.r, b.r),
			new(big.Rat).Mul(a.i, b.i)),
		new(big.Rat).Add(
			new(big.Rat).Mul(a.r, b.i),
			new(big.Rat).Mul(a.i, b.r))}
}

func div(a *ComplexR, b *ComplexR) *ComplexR {
	bunbo := new(big.Rat).Add(
		new(big.Rat).Mul(b.r, b.r),
		new(big.Rat).Mul(b.i, b.i))
	new_r := new(big.Rat).Quo(
		new(big.Rat).Add(
			new(big.Rat).Mul(a.r, b.r),
			new(big.Rat).Mul(a.i, b.i)),
		bunbo)
	new_i := new(big.Rat).Quo(
		new(big.Rat).Sub(
			new(big.Rat).Mul(b.r, a.i),
			new(big.Rat).Mul(a.r, b.i)),
		bunbo)
	return &ComplexR{new_r, new_i}
}

func pdist(a *ComplexR) *big.Rat {
	return new(big.Rat).Add(
		new(big.Rat).Mul(a.r, a.r),
		new(big.Rat).Mul(a.i, a.i))
}

func MainRat(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, newtonR(&ComplexR{
				new(big.Rat).SetFloat64(x),
				new(big.Rat).SetFloat64(y)}))
		}
	}
	png.Encode(out, img)
}

func newtonR(z *ComplexR) color.Color {
	const iterations = 37
	const contrast = 7
	if pdist(z).Cmp(new(big.Rat).SetFloat64(0)) == -1 {
		return color.Black
	}
	for i := uint8(0); i < iterations; i++ {
		z = minus(z, div(
			div(
				&ComplexR{
					new(big.Rat).SetFloat64(1),
					new(big.Rat).SetFloat64(0)},
				multi(z, multi(z, z))),
			&ComplexR{
				new(big.Rat).SetFloat64(4),
				new(big.Rat).SetFloat64(0)}))
		if pdist(minus(multi(z, multi(z, multi(z, z))), &ComplexR{
			new(big.Rat).SetFloat64(1),
			new(big.Rat).SetFloat64(0)})).Cmp(
			new(big.Rat).SetFloat64(1e-12)) == -1 {
			return color.RGBA{255 - contrast*i, 255, 0, 255}
		}
	}
	return color.Black
}
