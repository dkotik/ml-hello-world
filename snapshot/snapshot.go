package snapshot

import (
	"errors"
	"fmt"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func New(from *neuralnetwork.NeuralNetwork) (*Snapshot, error) {
	l := len(from.Layers)
	s := &Snapshot{
		Layers:  make([]uint16, l),
		Weights: make([]float64, 0, l*1024),
	}

	for _, layer := range from.Layers {
		for _, neuron := range layer.Neurons {
			for _, dendron := range neuron.Inbound {
				s.Weights = append(s.Weights, dendron.Weight)
			}
		}
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

type Snapshot struct {
	LearnedExamples uint16
	Layers          []uint16
	Weights         []float64
}

func (s *Snapshot) Validate() error {
	if s == nil {
		return errors.New("cannot use a <nil> neural network snapshot")
	}
	if len(s.Layers) < 2 {
		return errors.New("a neural network must contain at least two layers")
	}
	totalDendrons := 0
	neuronsInPreviousLayer := 0
	for i, neuronsInLayer := range s.Layers {
		if neuronsInLayer < 1 {
			return fmt.Errorf("network layer %d contains no Neurons", i+1)
		}
		totalDendrons += neuronsInPreviousLayer * int(neuronsInLayer)
	}

	if totalDendrons != len(s.Weights) {
		return fmt.Errorf("snapshot contains %d weight, but should contain %d weights instead", totalDendrons, len(s.Weights))
	}
	return nil
}
