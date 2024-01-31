package yamlutil

import (
	"os"

	"github.com/gofika/fileutil"
)

// ReadFileAny read struct from yml file
func ReadFileAny(filename string, v any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return ReadReaderAny(f, v)
}

// ReadFile read struct from yml file
func ReadFile[T any](filename string) (T, error) {
	var v T
	err := ReadFileAny(filename, &v)
	return v, err
}

// WriteFile write struct to yml file
func WriteFile(filename string, v any) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteWriter(f, v)
}

// WriteFileIndent write struct to yml file with indent
func WriteFileIndent(filename string, v any, spaces int) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return WriteWriterIndent(f, v, spaces)
}
