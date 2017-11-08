package capusta

import (
	"bytes"
	"encoding/binary"
)

func Binarizate(input ...interface{} ) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		binary.Write(buf, binary.LittleEndian, v)
	}
	return buf.Bytes()
}