package neuralnetwork

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

func EncodeFloat64(w io.Writer, f float64) error {
	return binary.Write(w, binary.LittleEndian, f)
}

func DecodeFloat64(f *float64, r io.Reader) error {
	return binary.Read(r, binary.LittleEndian, f)
}

type dumbReader struct {
	same []byte
}

func newDumbReaderFrom(f float64) *dumbReader {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return &dumbReader{same: buf[:]}
}

func (d *dumbReader) Read(b []byte) (n int, err error) {
	if len(b) != 8 {
		return 0, errors.New("float64 requires eight bytes")
	}
	copy(b, d.same)
	return 8, nil
}
