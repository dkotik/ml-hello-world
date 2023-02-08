package test

import (
	"testing"

	"github.com/moverest/mnist"
)

func TestLoadMNISTData(t *testing.T) {
	train, verify, err := mnist.Load("./data")
	if err != nil {
		t.Fatalf("could not load training data: %v", err)
	}

	t.Logf("found %d training images and %d labels", len(train.Images), len(train.Labels))
	t.Logf("found %d verify images and %d labels", len(verify.Images), len(verify.Labels))
}
