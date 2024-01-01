package ioutils

import (
	"bytes"
	"io"
)

// SniffReadCloser get data from stream and replace it with copy.
//
// Result is safe to modify and use in another go-routine.
func SniffReadCloser(stream *io.ReadCloser) ([]byte, error) {
	defer (*stream).Close()

	data, err := io.ReadAll(*stream)
	if err != nil {
		//nolint:wrapcheck
		return data, err
	}

	newData := make([]byte, len(data))
	copy(newData, data)

	*stream = io.NopCloser(bytes.NewReader(newData))

	return data, nil
}
