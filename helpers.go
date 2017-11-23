package capusta

import (
	"bytes"
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

