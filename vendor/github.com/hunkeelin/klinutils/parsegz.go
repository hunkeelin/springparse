package klinutils

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func ParseGz(g []byte) ([]byte, error) {
	ErrReturn := []byte("")
	r := bytes.NewReader(g)
	gz, err := gzip.NewReader(r)
	if err != nil {
		return ErrReturn, fmt.Errorf("Unable to read file in gzip %v", err)
	}
	defer gz.Close()
	buf := &bytes.Buffer{}
	buf.ReadFrom(gz)
	return buf.Bytes(), nil
}
