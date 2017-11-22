package capusta

import (
	"bytes"
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

func ConvertHashToString(input [32]byte) string {
	return hex.EncodeToString(input[:])
}
