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

	separator := []byte{}

	data := bytes.Join([][]byte{
		b.previousHash[:],
		[]byte(b.data),
		toBinary(b.timestamp),
		toBinary(proof),
	},
		separator,
	)

	return data

}

func (b *Block) ValidateHash() bool {
	h := sha256.New()
	h.Write(b.PrepareData(b.proof))
	hash := h.Sum(nil)
	return bytes.Equal(hash, []byte(b.hash[:]))
}
