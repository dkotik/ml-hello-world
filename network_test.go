package neuralnetwork

import "testing"

func TestNeuralNetworkWithoutLearning(t *testing.T) {
	nn, err := NewWithoutLearning(2, 2)
	if err != nil {
		t.Fatalf("failed to build a neural network: %v", err)
	}
	for i, l := range nn.Layers {
		t.Logf("- Layer %d", i+1)
		for j, n := range l.Neurons {
			t.Logf("  - Neuron %d.%d", i+1, j+1)
			for k, d := range n.Inbound {
				t.Logf("    - Dendron %d.%d.%d: %.2f", i+1, j+1, k+1, d.Strength)
				if d.Strength != 0.5 {
					t.Fatal("one dendron is not 0.5")
				}
			}
		}
	}
}
