package snapshot

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// FromReader builds a valid snapshot from an [io.Reader].
func FromReader(r io.Reader) (*Snapshot, error) {
	sn := &Snapshot{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&sn); err != nil {
		return nil, err
	}
	if err := sn.Validate(); err != nil {
		return nil, err
	}
	return sn, nil
}

func FromFile(p string) (*Snapshot, error) {
	handle, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("unable to load %q: %w", p, err)
	}
	defer handle.Close()
	return FromReader(handle)
}
