package utils

import (
	"encoding/binary"
	"io"
)

type byteReaderWrap struct {
	reader io.Reader
}

func (w *byteReaderWrap) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	_, err := w.reader.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], err
}

// Читает Varint из соединения
func ReadVarint(r io.Reader) (uint32, error) {
	v, err := binary.ReadUvarint(&byteReaderWrap{r})
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}
