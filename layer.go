package neuralnetwork

import (
	"fmt"
	"io"
)

type Layer struct {
	Neurons []*Neuron
}

func (l *Layer) ConnectTo(next *Layer, r io.Reader) (err error) {
	for _, n := range next.Neurons {
		if err = n.ConnectTo(l, r); err != nil {
			return fmt.Errorf("layer connection failed: %w", err)
		}
	}
	return nil
}

func NewLayer(size int) *Layer {
	neurons := make([]*Neuron, size)
	for i := 0; i < size; i++ {
		neurons[i] = &Neuron{}
	}

	return &Layer{
		Neurons: neurons,
	}
}
