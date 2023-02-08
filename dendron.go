package neuralnetwork

import (
	"encoding/binary"
	"io"
)

type Dendron struct {
	Strength float64
	Neuron   *Neuron
}

func (d *Dendron) Learn(desiredActivationStrength float64) {
	d.Strength = (d.Strength + desiredActivationStrength) / 2
}

func (d *Dendron) Encode(w io.Writer) error {
	return binary.Write(w, binary.LittleEndian, d.Strength)
}

func (d *Dendron) Decode(r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, &d.Strength)
}
