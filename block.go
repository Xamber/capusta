package capusta

import (
	"crypto/sha256"
)

type Block struct {
	index         int
	timestamp     int64
	data          transactions
	proof         int
	previousBlock *Block
}

func (b Block) hash() []byte {

	h := sha256.New()
	h.Write([]byte("hello world\n"))

	return h.Sum(nil)
}
