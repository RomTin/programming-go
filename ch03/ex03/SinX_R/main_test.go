package main

import (
	"math"
	"testing"
)

func TestIsValidPolygon(t *testing.T) {
	expected := false
	actual := isValidPolygon([]float64{1.0, 1.0, math.NaN()})
	if expected != actual {
		t.Errorf("Invalid point was passed through test.")
	}
}
