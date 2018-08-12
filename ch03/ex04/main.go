package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

// e.g.) http://localhost:8000/?color=FF00FF&size=0.5
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	color := r.FormValue("color")
	size, _ := strconv.ParseFloat(r.FormValue("size"), 64)
	draw(w, color, size)
}

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func draw(out io.Writer, color string, sc float64) {
	fmt.Fprint(out, fmt.Sprintf("<html><body><svg xmlms='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, sc)
			bx, by := corner(i, j, sc)
			cx, cy := corner(i, j+1, sc)
			dx, dy := corner(i+1, j+1, sc)
			if !(isValidPolygon([]float64{ax, ay, bx, by, cx, cy, dx, dy})) {
				continue
			}
			fmt.Fprint(out, fmt.Sprintf("<polygon style='stroke: #%s' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", color, ax, ay, bx, by, cx, cy, dx, dy))
		}
	}
	fmt.Fprint(out, "</svg></body></html>")
}

func isValidPolygon(arr []float64) bool {
	for _, p := range arr {
		if math.IsNaN(p) {
			return false
		}
	}
	return true
}

func corner(i, j int, sc float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y, sc)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y, sc float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r * sc
}
