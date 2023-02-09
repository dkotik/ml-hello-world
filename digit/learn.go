package digit

import (
	"fmt"
	"os"
)

func (d *DigitParser) Learn(digit uint8) {
	if d.target == digit {
		d.model.LearnFrom(d.model.LinearLearning(0.0, 1.0))
	} else {
		d.model.LearnFrom(d.model.LinearLearning(1.0, 0.0))
	}
}

func (d *DigitParser) Dump(p string) error {
	handle, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("cannot create a model file: %w", err)
	}
	defer handle.Close()
	return d.model.Dump(handle)
}
