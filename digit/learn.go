package digit

func (d *DigitParser) Learn(digit uint8) {
	if d.target == digit {
		d.lastLayer.Neurons[0].Learn(0)
		d.lastLayer.Neurons[1].Learn(1)
	} else {
		d.lastLayer.Neurons[0].Learn(1)
		d.lastLayer.Neurons[1].Learn(0)
	}
}
