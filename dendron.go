package neuralnetwork

type Dendron struct {
	Strength float64
	Neuron   *Neuron
}

func (d *Dendron) Learn(desiredActivationStrength float64) {
	d.Strength = (d.Strength + desiredActivationStrength) / 2
}
