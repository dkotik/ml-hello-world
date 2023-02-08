package neuralnetwork

type Dendron struct {
	Strength float64
	Neuron   *Neuron
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
