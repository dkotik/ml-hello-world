package digit

import (
	"errors"
	"fmt"
	"os"

	nn "github.com/dkotik/ml-hellow-world"
)

type DigitParser struct {
	target                uint8
	lastLayer             *nn.Layer
	recognizesNothing     *nn.Neuron
	recognizesTargetDigit *nn.Neuron
	model                 *nn.NeuralNetwork
}

func NewFromModel(digit uint8, model *nn.NeuralNetwork) (*DigitParser, error) {
	l := len(model.Layers)
	if l < 2 {
		return nil, errors.New("digit model must include at least two layers")
	}
	if len(model.Layers[l-1].Neurons) != 2 {
		return nil, errors.New("digit model must have exactly two neurons in the last layer")
	}

	result := &DigitParser{
		target: digit,
		model:  model,
	}

	l = l - 1
	result.lastLayer = model.Layers[l]
	result.recognizesNothing = model.Layers[l].Neurons[0]
	result.recognizesTargetDigit = model.Layers[l].Neurons[1]

	return result, nil
}

func New(digit uint8) (*DigitParser, error) {
	model, err := nn.NewWithoutLearning(28*28, 32, 16, 2)
	if err != nil {
		return nil, err
	}
	return NewFromModel(digit, model)
}

func NewFromFile(digit uint8, p string) (*DigitParser, error) {
	handle, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("cannot load data from %s: %w", p, err)
	}
	defer handle.Close()
	model, err := nn.New(handle, 28*28, 32, 16, 2)
	if err != nil {
		return nil, err
	}
	return NewFromModel(digit, model)
}

func Must(d *DigitParser, err error) *DigitParser {
	if err != nil {
		panic(err)
	}
	return d
}
