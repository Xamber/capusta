package capusta

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func isProofHash(hash [32]byte) bool {
	return bytes.HasPrefix(hash[:], defaultProof)
}

// Hashing make bytes buffer for Input argument and return sha256 hash
func Hashing(input interface{}) [32]byte {
	var encodingResult bytes.Buffer

	enc := gob.NewEncoder(&encodingResult)

	err := enc.Encode(input)
	handleError(err)

	hash := sha256.Sum256(encodingResult.Bytes())

	return hash
}

func ConvertHashToString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}
