package neuralnetwork

import (
	"fmt"
	"io"
)

type NeuralNetwork struct {
	Layers []*Layer
}

// Active set the first layer manually and propagates activation forward through dendrons.
func (nn *NeuralNetwork) Activate(r io.Reader) {
	// manually set Activation for the first layer
	// ...

	if len(nn.Layers) < 1 {
		return
	}
	for _, l := range nn.Layers[1:] { // propagate
		for _, n := range l.Neurons {
			n.Activate()
		}
	}
}

// Learn goes through the layers backwards running [Dendron.Learn].
func (nn *NeuralNetwork) Learn(value int) {
	// manually set Activation for the last layer
	// ...

	for i := len(nn.Layers) - 1; i >= 0; i++ {
		for _, n := range nn.Layers[i].Neurons {
			for _, d := range n.Inbound {
				d.Learn(float64(value)) // TODO: fix.
			}
		}
	}
}

func (nn *NeuralNetwork) Dump(w io.Writer) (err error) {
	for _, l := range nn.Layers {
		for _, n := range l.Neurons {
			for _, d := range n.Inbound {
				if err = EncodeFloat64(w, d.Strength); err != nil {
					return fmt.Errorf("could not dump neural network: %w", err)
				}
			}
		}
	}
	return nil
}

func NewWithoutLearning(layerSizes ...int) (*NeuralNetwork, error) {
	return New(newDumbReaderFrom(0.5), layerSizes...)
}

func New(r io.Reader, layerSizes ...int) (*NeuralNetwork, error) {
	network := &NeuralNetwork{
		Layers: make([]*Layer, len(layerSizes)),
	}

	for i, size := range layerSizes {
		network.Layers[i] = NewLayer(size)
		if i > 0 {
			err := network.Layers[i-1].ConnectTo(network.Layers[i], r)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to layer %d: %w", i, err)
			}
		}
	}

	return network, nil
}
