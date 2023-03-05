package minimum

import (
	"errors"
	"fmt"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func Must(nn neuralnetwork.MinimumFinder, err error) neuralnetwork.MinimumFinder {
	if err != nil {
		panic(err)
	}
	return nn
}

type options struct {
	minimumStep float64
	limit       int
}

func newOptions(withOptions []Option) (*options, error) {
	o := &options{}
	for _, option := range withOptions {
		if err := option(o); err != nil {
			return nil, err
		}
	}

	if o.minimumStep == 0 {
		o.minimumStep = 0.001
	}
	if o.limit == 0 {
		o.limit = 15
	}

	return o, nil
}

type Option func(*options) error

func WithAccuracy(step float64) Option {
	return func(o *options) error {
		if o.minimumStep != 0 {
			return errors.New("accuracy is already set")
		}
		if step <= 0 || step > 0.25 {
			return fmt.Errorf("cannot set step accuracy to %.2f", step)
		}
		o.minimumStep = step
		return nil
	}
}

func WithMaximumSearchDepthOf(steps int) Option {
	return func(o *options) error {
		if o.limit != 0 {
			return errors.New("maximum search depth is already set")
		}
		if steps <= 0 || steps > 1000 {
			return fmt.Errorf("cannot set maximum search depth to %d", steps)
		}
		o.limit = steps
		return nil
	}
}
