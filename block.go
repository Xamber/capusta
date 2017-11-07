package capusta

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	index         int
	timestamp     int64
	data          string
	proof         int
	previousBlock *Block
}

func (b Block) hash() string {

	h := sha256.New()
	h.Write([]byte(b.data))

	return hex.EncodeToString(h.Sum(nil))
}
