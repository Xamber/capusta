package capusta

import (
	"bytes"
	"encoding/binary"
	"log"
)

// Binarizate make bytes buffer for all input arguments and return all bytes
func Binarizate(input ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		err := binary.Write(buf, binary.LittleEndian, v)
		logError(err)
	}
	return buf.Bytes()
}

func logError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
