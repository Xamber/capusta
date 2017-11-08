package capusta

import (
	"crypto/sha256"
	"bytes"
)

type Block struct {
	index        int
	timestamp    int64
	data         string
	proof        int64
	hash         [32]byte
	previousHash [32]byte
}

func (b *Block) PrepareData(proof int64) []byte {
	data := Binarizate(b.previousHash, b.data, b.timestamp, proof)
	return data
}

func (b *Block) ValidateHash() bool {
	hash := sha256.Sum256(b.PrepareData(b.proof))
	return bytes.Equal(hash[:], []byte(b.hash[:]))
}
