package neuralnetwork

import (
	"fmt"
	"io"
	"math"
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
		value += math.Min(d.Strength, 0) * d.Neuron.Activation
	}
	n.Activation = value / float64(len(n.Inbound))
}

func (n *Neuron) Learn(desiredActivation float64) {
	// TODO: the magic happens below; it needs to be fixed
	n.Activation = (desiredActivation + n.Activation) / 2
	for _, d := range n.Inbound {
		if d.Strength > 0.5 && n.Activation < desiredActivation {
			d.Strength = math.Max(d.Strength+0.05, 1) // boost
		} else {
			d.Strength = math.Min(d.Strength-0.07, -1) // sink
		}
		d.Neuron.Learn(math.Min(d.Strength, 0) * desiredActivation)
	}

	// halfDelta := (desiredActivation - n.Activation) / 2
	// if halfDelta > 0 { // boost strength
	// 	for _, d := range n.Inbound {
	// 		if d.Strength > 0.5 {
	// 			d.Strength = math.Max(d.Strength+halfDelta, 1)
	// 		} else {
	// 			d.Strength = math.Min(d.Strength-halfDelta-0.05, 0)
	// 		}
	// 		d.Neuron.Learn(n.Activation + halfDelta)
	// 	}
	// } else { // lower strength
	// 	for _, d := range n.Inbound {
	// 		if d.Strength > 0.5 {
	// 			d.Strength = math.Min(d.Strength+halfDelta*2, 0)
	// 		} else {
	// 			d.Strength = math.Max(d.Strength-halfDelta, 1)
	// 		}
	// 		d.Neuron.Learn(n.Activation + halfDelta)
	// 	}
	// }
	// 	// activation(d.Strength + desiredActivationStrength) / 2
	// 	// d.Learn(desiredActivationStrength)
	// 	d.Neuron.Learn(desiredActivationStrength * d.Strength) // propagate
	// }
}
