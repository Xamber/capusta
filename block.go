package capusta

import (
	"crypto/sha256"
	"bytes"
	"fmt"
)

// block contain information about block
type block struct {
	index        int
	timestamp    int64
	data         string
	proof        int64
	hash         [32]byte
	previousHash [32]byte
}

// block.prepareData create binary slice from block and found proof
func (b *block) prepareData(proof int64) []byte {
	return Binarizate(b.previousHash, b.data, b.timestamp, proof)
}

// block.validate check hash of block
func (b *block) validate() bool {
	hash := sha256.Sum256(b.prepareData(b.proof))
	return bytes.HasPrefix(hash[:], b.hash[:])
}

// block.info return string with info about block
func (b *block) info() string {
	template := "block Index: %v Timestamp: %v Proof: %v \nHash: %x\nPreviousHash: %x\nValidated: %v\n"
	return fmt.Sprintf(template, b.index, b.timestamp, b.proof, b.hash, b.previousHash, b.validate())
}
