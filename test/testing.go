package test

import (
	"fmt"

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
