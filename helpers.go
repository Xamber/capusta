package capusta

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/gob"
	"crypto/sha256"
	"encoding/hex"
)

// Binarizate make bytes buffer for all Input arguments and return all bytes
func Binarizate(input ...interface{}) []byte {
	buf := new(bytes.Buffer)
	for _, v := range input {
		err := binary.Write(buf, binary.LittleEndian, v)
		handleError(err)
	}
	return buf.Bytes()
}

// Hashing make bytes buffer for Input argument and return sha256 hash
func Hashing(input interface{}) [32]byte {
	encoded := new(bytes.Buffer)

	enc := gob.NewEncoder(encoded)

	err := enc.Encode(input)
	handleError(err)

	hash := sha256.Sum256(encoded.Bytes())

	return hash
}

func ConvertHashToString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
