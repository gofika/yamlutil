package yamlutil

import (
	"io"

	"gopkg.in/yaml.v3"
)

// ReadReaderAny read struct from yml reader
func ReadReaderAny(r io.Reader, v any) error {
	dec := yaml.NewDecoder(r)
	err := dec.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// ReadReader read struct from yml reader
func ReadReader[T any](r io.Reader) (T, error) {
	var v T
	err := ReadReaderAny(r, &v)
	return v, err
}
