package snapshot

import (
	"bytes"
	"testing"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func TestWrite(t *testing.T) {
	b := &bytes.Buffer{}

	nn, err := neuralnetwork.NewWithoutLearning(28*28, 32, 16, 2)
	if err != nil {
		t.Fatal("cannot create a neural network:", err)
	}

	s1, err := New(nn)
	if err != nil {
		t.Fatal("cannot create snapshot:", err)
	}

	if err = s1.WriteTo(b); err != nil {
		t.Fatal("cannot write:", err)
	}

	s2, err := FromReader(b)
	if err != nil {
		t.Fatal("cannot read:", err)
	}

	t.Fatal("should compare s1 to s2", s1, s2)
}
