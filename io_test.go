package neuralnetwork

import "testing"

func TestByteToFloat64(t *testing.T) {
	cases := []struct {
		In  byte
		Out float64
	}{
		{0x00, 0},
		{0x63, 0.38823529411},
		{0x80, 0.5},
		{0xFF, 1},
	}

	for _, tcase := range cases {
		fl := ByteToFloat64(tcase.In)
		if delta := tcase.Out - fl; delta < -0.002 || delta > 0.002 {
			t.Fatalf("Values did not match: %.8f vs %.8f", fl, tcase.Out)
		}
	}
}
