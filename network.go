package neuralnetwork

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

type NeuralNetwork struct {
	Layers []*Layer
}

// Active set the first layer manually and propagates activation forward through dendrons.
func (nn *NeuralNetwork) Activate(r io.Reader) (err error) {
	if len(nn.Layers) <= 1 {
		return errors.New("action requires at least two layers")
	}

	// manually set Activation for the first layer
	bufferedReader := bufio.NewReader(r)
	var (
		i   int
		one byte
	)
	for ; i < len(nn.Layers[0].Neurons); i++ {
		one, err = bufferedReader.ReadByte()
		if err != nil {
			return err
		}
		nn.Layers[0].Neurons[i].Activation = ByteToFloat64(one)
		nn.Layers[0].Neurons[i].Propagate()
	}

	// for _, l := range nn.Layers[1:] { // propagate
	// 	for _, n := range l.Neurons {
	// 		n.Activate()
	// 	}
	// }

	return nil
}

// // LearnFrom goes through the layers backwards running [Dendron.Learn].
// func (nn *NeuralNetwork) LearnFrom(f DeviancyMeasure) {
// 	if len(nn.Layers) <= 1 {
// 		panic("action requires at least two layers")
// 	}
//
// 	for _, n := range nn.Layers[len(nn.Layers)-1].Neurons {
// 		for _, d := range n.Inbound {
// 			d.LearnFrom(f)
// 		}
// 	}
// }

func (nn *NeuralNetwork) Dump(w io.Writer) (err error) {
	for _, l := range nn.Layers {
		for _, n := range l.Neurons {
			for _, d := range n.Inbound {
				if err = EncodeFloat64(w, d.Weight); err != nil {
					return fmt.Errorf("could not dump neural network: %w", err)
				}
			}
		}
	}
	return nil
}

func NewWithoutLearning(layerSizes ...int) (*NeuralNetwork, error) {
	return New(newDumbReaderFrom(1.0), layerSizes...)
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
