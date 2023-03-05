package minimum

import (
	"math"
	"testing"
)

func TestDerivative(t *testing.T) {
	cases := []struct {
		x1, y1 float64
		x2, y2 float64
		result float64
	}{
		{0, 0, 0, 0, math.MaxFloat64},                 // vertical up
		{1, 1, 2, 1, 0},                               // flat
		{1, 1, 2, 2, 1},                               // slope up
		{2, 2, 1, 1, 1},                               // slope down
		{1, 1, 2, math.MaxFloat64, math.MaxFloat64},   // vertical up
		{0, 0, 1, -math.MaxFloat64, -math.MaxFloat64}, // vertical down
		{-1, -1, 2, 2, 1},                             // slope up
		{2, 2, -1, -1, 1},                             // slope up
		{0, 0, 2, -2, -1},                             // slope down
		{0, 0, 1, -1, -1},                             // slope down
	}

	for i, c := range cases {
		if result := Derivative(c.x1, c.y1, c.x2, c.y2); math.Abs(result-c.result) > 0.000001 {
			t.Errorf("expected derivative value for case %d did not match: %.6f vs %.6f", i, result, c.result)
		}
	}
}
