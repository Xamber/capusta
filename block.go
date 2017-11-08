package capusta

import (
	"crypto/sha256"
	"bytes"
	"fmt"
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

func (b *Block) Validate() bool {
	hash := sha256.Sum256(b.PrepareData(b.proof))
	return bytes.Equal(hash[:], []byte(b.hash[:]))
}

func (b *Block) Info() string {
	template := `Block Index: %v Time: %d
Hash: %x
PreviousHash: %x
Validated: %v
`
	return fmt.Sprintf(template, b.index, b.timestamp, b.hash, b.previousHash, b.Validate())
}
