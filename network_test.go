package neuralnetwork

import "testing"

func TestNeuralNetworkWithoutLearning(t *testing.T) {
	nn, err := New()
	if err != nil {
		t.Fatalf("failed to build a neural network: %v", err)
	}
	for i, l := range nn.Layers {
		t.Logf("- Layer %d", i+1)
		for j, n := range l.Neurons {
			t.Logf("  - Neuron %d.%d", i+1, j+1)
			for k, d := range n.Inbound {
				t.Logf("    - Dendron %d.%d.%d: %.2f", i+1, j+1, k+1, d.Weight)
				if d.Weight == 0 {
					t.Fatal("one dendron is not set")
				}
			}
		}
	}
}
