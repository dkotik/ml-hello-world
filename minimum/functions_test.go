package minimum

import (
	"math"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func closeEnoughTo(delta, x1, x2 float64) bool {
	return math.Abs(x2-x1) < delta
}

var testCases = []*struct {
	Name            string
	Function        neuralnetwork.Loss
	ExpectedMinimum float64
}{
	{
		Name: "(x-0.5)^2+0.5",
		Function: func(x1 float64) float64 {
			return math.Pow(x1-0.5, 2) + 0.5
		},
		ExpectedMinimum: 0.5,
	},
	{
		Name: "(2x)^3+(2x-0.7)^2+0.2",
		Function: func(x1 float64) float64 {
			return math.Pow(x1*2, 3) + math.Pow(x1*2-0.7, 2) + 0.2
		},
		ExpectedMinimum: 0.213,
	},
}
