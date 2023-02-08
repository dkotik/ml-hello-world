package digit

import (
	"bytes"
	"log"
)

func (d *DigitParser) Load(image []byte) bool {
	if err := d.model.Activate(bytes.NewReader(image)); err != nil {
		panic(err)
	}
	isTargetDigit := d.recognizesTargetDigit.Activation
	isNothing := d.recognizesNothing.Activation

	if isTargetDigit > isNothing+.05 {
		log.Printf(
			"recognized %d, %d%% certain, %d%% suspect that predication is wrong",
			d.target,
			int(isTargetDigit*100),
			int(isNothing*100),
		)
		return true
	}
	return false
}
