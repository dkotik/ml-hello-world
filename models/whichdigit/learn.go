package whichdigit

import (
	"math"

	"github.com/dkotik/ml-hellow-world/minimum"
)

func (d *Detector) Learn(image []byte, label Digit) (err error) {
	if err = d.IsDetectableImage(image); err != nil {
		return err
	}
	d.activateUsingImage(image)

	d.network.Learn(func() (score float64) {
		for digit, neuron := range d.digits {
			if digit == label {
				score += math.Abs(1.0 - neuron.Activation)
				continue
			}
			score += math.Abs(0.0 - neuron.Activation)
		}
		return score
	}, minimum.NewLinear())
	return nil
}
