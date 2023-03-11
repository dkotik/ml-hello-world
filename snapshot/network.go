package snapshot

import (
	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func (s *Snapshot) NeuralNetwork() (*neuralnetwork.NeuralNetwork, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}

	processedDendrons := 0
	return &neuralnetwork.New(
		neuralnetwork.WithLayers(s.Layers...),
		neuralnetwork.WithDendronFactory(
			func(n *neuralnetwork.Neuron) (*neuralnetwork.Dendron, error) {
				d := &neuralnetwork.Dendron{
					Weight: s.Weights[d],
					Neuron: n,
				}
				processedDendrons++
				return d, nil
			},
		),
	), nil
}
