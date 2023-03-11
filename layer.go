package neuralnetwork

import "fmt"

type Layer struct {
	Neurons []*Neuron
}

func (l *Layer) ConnectTo(next *Layer, using DendronFactory) (err error) {
	for _, current := range l.Neurons {
		current.NextLayer = next
	}

	for _, next := range next.Neurons {
		next.Inbound = make([]*Dendron, len(l.Neurons))
		for i, inboundNeuron := range l.Neurons {
			next.Inbound[i], err = using(inboundNeuron)
			if err != nil {
				return fmt.Errorf("cannot create a dendron: %w", err)
			}
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
