package capusta

import (
	"bytes"
	"encoding/binary"
)

// Binarizate make bytes buffer for all input arguments and return all bytes
func Binarizate(input ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		binary.Write(buf, binary.LittleEndian, v)
	}
	return buf.Bytes()
}
