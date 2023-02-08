package neuralnetwork

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
)

type NeuralNetwork struct {
	Layers []*Layer
}

type dumbReader struct {
	same []byte
}

func (d *dumbReader) Read(b []byte) (n int, err error) {
	if len(b) != 8 {
		return 0, errors.New("float64 requires eight bytes")
	}
	// copy(b, float64(0.5))
	copy(b, d.same)
	return 8, nil
}

func NewWithoutLearning(layerSizes ...int) (*NeuralNetwork, error) {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(0.5))

	return New(&dumbReader{same: buf[:]}, layerSizes...)
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
