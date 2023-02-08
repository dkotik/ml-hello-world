package neuralnetwork

import (
	"fmt"
	"io"
)

type Neuron struct {
	Activation float64
	Inbound    []*Dendron
	// Outbound []*Dendron
}

func (n *Neuron) ConnectTo(l *Layer, r io.Reader) (err error) {
	n.Inbound = make([]*Dendron, len(l.Neurons))
	for i, inboundNeuron := range l.Neurons {
		d := &Dendron{
			Neuron: inboundNeuron,
		}
		if err = DecodeFloat64(&d.Strength, r); err != nil {
			return fmt.Errorf("could not decode dendron data: %w", err)
		}
		n.Inbound[i] = d
	}
	return nil
}

func (n *Neuron) Activate() {
	value := float64(0)
	for _, d := range n.Inbound {
		value += d.Strength * d.Neuron.Activation
	}
	n.Activation = value / float64(len(n.Inbound))
}

func (n *Neuron) Learn(desiredActivationStrength float64) {
	for _, d := range n.Inbound {
		d.Learn(desiredActivationStrength)
		d.Neuron.Learn(desiredActivationStrength) // propagate
	}
}
