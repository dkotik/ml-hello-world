package minimum

import (
	"fmt"
	"log"
	"math"

	neuralnetwork "github.com/dkotik/ml-hellow-world"
)

func NewLinear(withOptions ...Option) (neuralnetwork.MinimumFinder, error) {
	options, err := newOptions(withOptions)
	if err != nil {
		return nil, fmt.Errorf("cannot create a linear minimum finder: %w", err)
	}
	return &linear{
		step:        options.minimumStep,
		minimumStep: options.minimumStep,
		limit:       options.limit,
	}, nil
}

type linear struct {
	step        float64
	minimumStep float64
	limit       int
}

// FindMinimum finds a value that returns the minimum [neuralnetwork.Loss].
func (d *linear) FindMinimum(x1 float64, f neuralnetwork.Loss) (best float64) {
	var (
		step       = d.step
		x2, y1, y2 float64
		lastFlip   float64
		i          int
	)
	best = x1

	for i = 0; i < d.limit; i++ {
		y1 = f(x1)
		x2 = x1 + step

		if x2 > 1.0 {
			x2 = 1.0
		} else if x2 < 0 {
			x2 = 0
		}
		y2 = f(x2)

		if y1 <= y2 { // colder
			if math.Abs(step) < d.minimumStep {
				step *= -1          // flip direction
				if lastFlip == x1 { // detect double consequitive flip
					// log.Println("end", y1, y2, x1)
					break
				}
				lastFlip = x1
				// log.Println("flip", y1, y2, x1)
				continue
			}
			// log.Println("colder", y1, y2, x2)
			step *= 0.4
			continue
		}
		// warmer
		// log.Println("warmer", y1, y2, x2)
		best = x2
		x1 = x2
		step *= 2
	}

	d.step = step
	log.Printf("found %.6f value after %d attempts", x2, i)
	return best
}
