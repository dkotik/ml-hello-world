package test

import (
	"fmt"
	"testing"

	"github.com/dkotik/ml-hellow-world/digit"
	"github.com/moverest/mnist"
)

const targetRecognitionDigit = 1

func TestDigitRecognition(t *testing.T) {
	train, verify, err := mnist.Load("./data")
	if err != nil {
		t.Fatalf("could not load training data: %v", err)
	}

	trainL := len(train.Images)
	if trainL == 0 {
		t.Fatal("no training images were loaded")
	}

	nn := digit.Must(digit.New(targetRecognitionDigit))

	for i := 0; i < len(train.Images) && i < 999999; i++ {
		if nn.Load(train.Images[i][:]) {
			dumpImage(train.Images[i])
		}
		nn.Learn(uint8(train.Labels[i]))
		// _, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		//
	}

	verifyL := len(verify.Images)
	t.Run(fmt.Sprintf("verify against %d images", verifyL), func(t *testing.T) {
		matched := 0
		possible := 0
		for i := 0; i < len(verify.Images) && i < 999999; i++ {
			if train.Labels[i] != targetRecognitionDigit {
				continue // skip other digits
			}
			possible++
			if nn.Load(verify.Images[i][:]) {
				matched++
			}
		}
		t.Logf("Model regonized %d images and has %.6f%% recognition strength", matched, float32(matched*100)/float32(possible))
	})

}
