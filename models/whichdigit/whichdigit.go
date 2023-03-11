package whichdigit

import (
	neuralnetwork "github.com/dkotik/ml-hellow-world"

	"fmt"
)

type Detector struct {
	network        *neuralnetwork.NeuralNetwork
	firstLayerSize int
	firstLayer     *neuralnetwork.Layer
	digits         map[Digit]*neuralnetwork.Neuron
}

func New(withOptions ...Option) (*Detector, error) {
	o := &options{}
	var err error
	for _, option := range append(
		withOptions,
		WithDefaultOptions(),
		func(o *options) error {
			neuronCount := len(o.NeuralNetwork.Layers[len(o.NeuralNetwork.Layers)-1].Neurons)
			digitCount := len(o.Digits)
			if neuronCount != digitCount {
				return fmt.Errorf("the number of neurons in the last layer does not match the number of digits: %d vs %d", neuronCount, digitCount)
			}
			return nil
		},
	) {
		if err = option(o); err != nil {
			return nil, fmt.Errorf("could not create a digit detector: %w", err)
		}
	}

	digits := make(map[Digit]*neuralnetwork.Neuron)
	lastLayer := o.NeuralNetwork.Layers[len(o.NeuralNetwork.Layers)-1].Neurons
	for i, digit := range o.Digits {
		digits[digit] = lastLayer[i]
	}

	return &Detector{
		network:        o.NeuralNetwork,
		firstLayerSize: len(o.NeuralNetwork.Layers[0].Neurons),
		firstLayer:     o.NeuralNetwork.Layers[0],
		digits:         digits,
	}, nil
}
