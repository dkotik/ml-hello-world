package neuralnetwork

import "errors"

type options struct {
	DendronFactory DendronFactory
	Layers         []int
}

type Option func(*options) error

func WithDefaultOptions() Option {
	return func(o *options) error {
		if o.DendronFactory == nil {
			o.DendronFactory = NewPseudoRandomDendronFactory(0.1)
		}
		if len(o.Layers) == 0 {
			o.Layers = []int{4, 4, 2}
		}
		return nil
	}
}

func WithLayers(sizes ...uint16) Option {
	return func(o *options) error {
		if o.Layers != nil {
			return errors.New("layer sizes have already been set")
		}
		if len(sizes) < 2 {
			return errors.New("a neural network must have at least two layers")
		}
		o.Layers = make([]int, len(sizes))
		for i, size := range sizes {
			if size < 1 {
				return errors.New("cannot use a layer with fewer than 1 neuron")
			}
			o.Layers[i] = int(size)
		}
		return nil
	}
}

func WithDendronFactory(f DendronFactory) Option {
	return func(o *options) error {
		if f == nil {
			return errors.New("cannot use <nil> dendron factory")
		}
		if o.DendronFactory != nil {
			return errors.New("dendron factory is already set")
		}
		o.DendronFactory = f
		return nil
	}
}
