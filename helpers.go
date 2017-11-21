package capusta

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/gob"
	"crypto/sha256"
)

// Binarizate make bytes buffer for all Input arguments and return all bytes
func Binarizate(input ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		err := binary.Write(buf, binary.LittleEndian, v)
		logError(err)
	}
	return buf.Bytes()
}

func Hashing(input interface{}) [32]byte {
	encoded := new(bytes.Buffer)

	enc := gob.NewEncoder(encoded)
	err := enc.Encode(input)

	logError(err)

	hash := sha256.Sum256(encoded.Bytes())

	return hash
}

func logError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
