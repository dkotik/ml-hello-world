package neuralnetwork

type Dendron struct {
	Strength float64
	Neuron   *Neuron
}

func (d *Dendron) LearnFrom(f DeviancyMeasure) {
	bestDeviancy := f()
	// bestStrength := d.Strength
	step := -0.05
	for i := 0; i < 10; i++ {
		d.Strength += step
		if d.Strength < 0 {
			step = step * -1 * 0.5 // flip direction and reduce magntitude
			d.Strength = 0
			continue
		} else if d.Strength > 1 {
			step = step * -1 * 0.5 // flip direction and reduce magntitude
			d.Strength = 1
			continue
		}

		if newDeviancy := f(); newDeviancy < bestDeviancy {
			bestDeviancy = newDeviancy
			// bestStrength = d.Strength
		} else {
			step = step * -1 * 0.5 // flip direction and reduce magntitude
		}
	}

	for _, child := range d.Neuron.Inbound {
		child.LearnFrom(f)
	}
}

// func (d *Dendron) Learn(desiredActivationStrength float64) {
// 	d.Strength = (d.Strength + desiredActivationStrength) / 2
// 	d.Neuron.Learn(d.Strength)
// 	// delta := d.Strength - desiredActivationStrength
// 	// d.Strength = d.Strength + delta*2
// 	//
// 	// if d.Strength < 0 {
// 	// 	d.Strength = 0
// 	// } else if d.Strength > 1 {
// 	// 	d.Strength = 1
// 	// }
// }
