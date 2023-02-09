package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dkotik/ml-hellow-world/digit"
	"github.com/moverest/mnist"
)

var offset = flag.Int("offset", 0, "which image index to begin from")
var batch = flag.Int("limit", 10, "how many images to consume")
var output = flag.String("output", "digit-1-model.bin", "where to save the model")
var data = flag.String("data", "../data", "where to load training data from")

func load(targetDigit uint8) *digit.DigitParser {
	model, err := digit.NewFromFile(targetDigit, *output)
	if err != nil {
		return digit.Must(digit.New(targetDigit))
	}
	return model
}

func trainUsingData(model *digit.DigitParser) error {
	train, _, err := mnist.Load(*data)
	if err != nil {
		return fmt.Errorf("could not load training data: %w", err)
	}

	for i := *offset; i < *offset+*batch || i < len(train.Images); i++ {
		model.Load(train.Images[i][:])
		model.Learn(uint8(train.Labels[i]))

		if i%100 == 0 {
			log.Printf("Processed image #%d.", i)
		}
	}
	model.Dump(*output)

	return nil
}

func main() {
	flag.Parse()

	model := load(1)
	if err := trainUsingData(model); err != nil {
		log.Fatalln("could not train data: ", err)
	}
}
