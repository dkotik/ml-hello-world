package test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/moverest/mnist"
)

func TestLoadMNISTData(t *testing.T) {
	t.Skip("checked, does work")

	train, verify, err := mnist.Load("./data")
	if err != nil {
		t.Fatalf("could not load training data: %v", err)
	}

	trainL := len(train.Images)
	if trainL == 0 {
		t.Fatal("no training images were loaded")
	}

	t.Logf("found %d training images and %d labels", trainL, len(train.Labels))
	t.Logf("found %d verify images and %d labels", len(verify.Images), len(verify.Labels))

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(trainL)
	t.Logf("showing a random image %d:", i)
	dumpImage(
		train.Images[i],
	)
}
