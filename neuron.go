package neuralnetwork

import (
	"fmt"
	"io"
)

type Neuron struct {
	Activation float64
	Inbound    []*Dendron

	// NextLayer is used for updating output neurons iteratively to calculate Loss faster instead of going through every single [Neuron] in a [NeuralNetwork].
	NextLayer *Layer
	// Outbound   []*Dendron
}

func (n *Neuron) ConnectTo(l *Layer, r io.Reader) (err error) {
	n.Inbound = make([]*Dendron, len(l.Neurons))
	for i, inboundNeuron := range l.Neurons {
		d := &Dendron{
			Neuron: inboundNeuron,
		}
		if err = DecodeFloat64(&d.Weight, r); err != nil {
			return fmt.Errorf("could not decode dendron data: %w", err)
		}
		n.Inbound[i] = d
		inboundNeuron.NextLayer = l
		// inboundNeuron.Outbound = append(
		// 	inboundNeuron.Outbound, n,
		// )
	}
	return nil
}

func (n *Neuron) Activate() {
	value := float64(0)
	for _, d := range n.Inbound {
		value += d.Weight * d.Neuron.Activation
	}
	n.Activation = value / float64(len(n.Inbound))
}

func (n *Neuron) Propagate() {
	for _, outboundNeuron := range n.NextLayer.Neurons {
		outboundNeuron.Activate()
		n.Propagate()
	}
}
