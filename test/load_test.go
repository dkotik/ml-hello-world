package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/moverest/mnist"
)

func dumpImage(image *mnist.Image) {
	for i := 0; i < 28; i++ {
		line := [28]byte{}
		for j := 0; j < 28; j++ {
			if image[i*28+j] > 127 {
				line[j] = '*'
			} else {
				line[j] = ' '
			}
		}
		fmt.Printf("%s\n", line[:])
	}
}

func TestLoadMNISTData(t *testing.T) {
	t.Skip() // does work

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
