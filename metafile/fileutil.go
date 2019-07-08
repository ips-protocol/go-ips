package metafile

import (
	"bytes"
	"io"
)

func GetMeta(data []byte) (*Meta, error) {
	return DecodeMeta(data)
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
