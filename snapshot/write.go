package snapshot

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func (s *Snapshot) WriteTo(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(s)
}

func (s *Snapshot) WriteToFile(p string) error {
	handle, err := os.Create(p)
	if err != nil {
		return fmt.Errorf("unable to save to file %q: %w", p, err)
	}
	defer handle.Close()
	return s.WriteTo(handle)
}
