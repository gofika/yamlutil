package yamlutil

import (
	"io"

	"gopkg.in/yaml.v3"
)

// WriteWriter write struct to yml writer
func WriteWriter(w io.Writer, v any) error {
	enc := yaml.NewEncoder(w)
	return enc.Encode(v)
}

// WriteWriterIndent write struct to yml writer with indent
func WriteWriterIndent(w io.Writer, v any, spaces int) error {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(spaces)
	return enc.Encode(v)
}
