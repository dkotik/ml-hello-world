package neuralnetwork

func (nn *NeuralNetwork) Learn(loss Loss, finder MinimumFinder) {
	nn.LearnedSamplesCount++
	var startingWeight, betterWeight float64
	for _, layer := range from.Layers {
		for _, neuron := range layer.Neurons {
			for _, dendron := range neuron.Inbound {
				startingWeight = dendron.Weight
				betterWeight = finder.FindMinimum(
					startingWeight,
					func(tryWeight float64) (score float64) {
						dendron.Weight = tryWeight // adjust
						dendron.Neuron.Activate()  // activate
						dendon.Neuron.Propagate()  // propagate
						return loss()              // calculate loss
					},
				)
				dendron.Weight = startingWeight + ((betterWeight - startingWeight) / nn.LearnedSamplesCount)
			}
		}
	}
}
