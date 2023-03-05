package minimum

import "math"

// Derivative returns the slope of the function between two coordinates. Zero value indicates a flat section. Large value indicates a vertical fall. Negative values indicate down slope.
func Derivative(x1, y1, x2, y2 float64) float64 {
	dx := (x2 - x1)
	dy := (y2 - y1)
	if dx == 0 {
		if dy >= 0 {
			return math.MaxFloat64
		}
		return -math.MaxFloat32
	}
	return dy / dx
}
