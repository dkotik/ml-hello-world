package neuralnetwork

import (
	"errors"
	"math"
)

func (nn *NeuralNetwork) LinearLoss(values ...float64) (Loss, error) {
	lastLayer := nn.Layers[len(nn.Layers)-1]
	if len(lastLayer.Neurons) != len(values) {
		return errors.New("linear learning requires the same number of values as the number of neurons in the last layer of the neural network")
	}

	return func() (score float64) {
		for i := 0; i < len(values); i++ {
			score += math.Abs(values[i] - lastLayer.Neurons[i].Activation)
		}
		return score
	}, nil
}

// Loss measures the deviancy between the model and the learning data. [NeuralNetwork] learns by searching for a minimum loss value using a [GlenBottomFinder].
type Loss func() float64

type ParameterizedLoss func(float64) float64

// MinimumFinder finds the value between 0.0 and 1.0 that yields the lowest return value of a function.
type MinimumFinder interface {
	FindMinimum(start float64, f ParameterizedLoss) float64
}

// DeviancyMeasure is used by [Dendron] to find a [Dendron.Strength] value that brings the deviancy closest towards 0.
type DeviancyMeasure func() (score float64)

func (nn *NeuralNetwork) LinearLearning(values ...float64) DeviancyMeasure {
	if len(nn.Layers) < 2 {
		panic("a neural network must have at least two layers")
	}
	lastLayer := nn.Layers[len(nn.Layers)-1]
	if len(lastLayer.Neurons) != len(values) {
		panic("linear learning requires the same number of values as the number of neurons in the last layer of the neural network")
	}

	return func() (score float64) {
		for i := 0; i < len(values); i++ {
			score += math.Abs(values[i] - lastLayer.Neurons[i].Activation)
		}
		return score
	}
}
