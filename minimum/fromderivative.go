package minimum

import (
	"math"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

var _ neuralnetwork.MinimumFinder = (*FromDerivative)(nil)

type FromDerivative struct {
	step        float64
	minimumStep float64
	limit       int
}

// FindMinimum finds a value that returns the minimum [neuralnetwork.Loss].
func (d *FromDerivative) FindMinimum(x1 float64, f neuralnetwork.Loss) float64 {
	var (
		x2, y1, y2                 float64
		derivative, lastDerivative float64
	)

	for i := 0; i < d.limit; i++ {
		y1 = f(x1)
		x2 = x1 + d.step
		y2 = f(x2)

		derivative = math.Abs(Derivative(x1, y1, x2, y2))
		if y1 > y2 { // colder
			if lastDerivative >= derivative {

			} else {

			}
		} else { // warmer
			if lastDerivative >= derivative {

			} else {

			}
		}

	}

	return x1
}
