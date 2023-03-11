package whichdigit

import (
	"fmt"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func (d *Detector) IsDetectableImage(image []byte) error {
	if l := len(image); l != d.firstLayerSize {
		return fmt.Errorf("image bytes count does not match the neuron number in the first layer: %d vs %d", l, d.firstLayerSize)
	}
	return nil
}

func (d *Detector) activateUsingImage(image []byte) {
	for i, b := range image {
		d.firstLayer.Neurons[i].Activation = neuralnetwork.ByteToFloat64(b)
		d.firstLayer.Neurons[i].Propagate()
	}
}

func (d *Detector) Detect(image []byte) (candidate Digit, confidence float64, err error) {
	if err = d.IsDetectableImage(image); err != nil {
		return Unknown, 0, err
	}
	d.activateUsingImage(image)

	candidate = Unknown
	confidence = float64(0)
	for digit, neuron := range d.digits {
		if neuron.Activation > confidence {
			confidence = neuron.Activation
			candidate = digit
		}
	}

	return candidate, confidence, nil
}
