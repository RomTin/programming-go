package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	min, max := calcborder()
	fmt.Printf("<svg xml='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if !(isValidPolygon([]float64{ax, ay, bx, by, cx, cy, dx, dy})) {
				continue
			}
			fmt.Printf("<polygon style='stroke: %s' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", getcolor(i, j, min, max), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func isValidPolygon(arr []float64) bool {
	for _, p := range arr {
		if math.IsNaN(p) {
			return false
		}
	}
	return true
}

func corner(i, j int) (float64, float64) {
	x, y := cv(i, j)
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func cv(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return x, y
}

func calcborder() (float64, float64) {
	min, max := float64(0), float64(0)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x, y := cv(i, j)
			z := f(x, y)
			if z > max {
				max = z
			}
			if z < min {
				min = z
			}
		}
	}
	return min, max
}

func getcolor(i, j int, min, max float64) string {
	mid := (max + min) / 2
	dist := (max - min) / 2
	x, y := cv(i, j)
	z := f(x, y)
	red, blue := 0, 0
	if z > mid {
		red = 255
		blue = int((z - mid) / dist * 255)
	} else {
		red = 255 - int((mid-z)/dist*255)
		blue = 255
	}
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
