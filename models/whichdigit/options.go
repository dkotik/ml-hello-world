package whichdigit

import (
	neuralnetwork "github.com/dkotik/ml-hellow-world"

	"errors"
	"fmt"
)

type options struct {
	Digits        []Digit
	NeuralNetwork *neuralnetwork.NeuralNetwork
}

type Option func(*options) error

func WithDefaultOptions() Option {
	return func(o *options) (err error) {
		defer func() {
			if err != nil {
				err = fmt.Errorf("could not apply default option: %w", err)
			}
		}()

		if o.Digits == nil {
			if err = WithDigits(Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine)(o); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithDigits(ds ...Digit) Option {
	return func(o *options) error {
		if o.Digits != nil {
			return errors.New("digits have already been selected")
		}
		duplicates := make(map[Digit]struct{})
		result := make([]Digit, len(ds))
		for i, d := range ds {
			if d < Unknown || d > Nine {
				return fmt.Errorf("cannot use digit %d for recognition", d)
			}
			if _, ok := duplicates[d]; ok {
				return fmt.Errorf("the same digit %d cannot be used twice", d)
			}
			duplicates[d] = struct{}{}
			result[i] = d
		}
		o.Digits = result
		return nil
	}
}

func WithNeuralNetwork(nn *neuralnetwork.NeuralNetwork) Option {
	return func(o *options) error {
		if o.NeuralNetwork != nil {
			return errors.New("neural network is already loaded")
		}
		if nn == nil {
			return errors.New("cannot use a <nil> neural network")
		}
		if len(nn.Layers) < 2 {
			return errors.New("cannot use a neural network with less than two layers")
		}
		o.NeuralNetwork = nn
		return nil
	}
}
