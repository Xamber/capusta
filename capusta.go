package capusta

import (
	"time"
	"crypto/sha256"
)

const DEFAULT_PROOF string = "000"

var Blockchain blockchain

func init() {

	genesisBlock := Block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         "",
		proof:        1337,
		previousHash: []byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}

func Hash(data string) []byte {

	h := sha256.New()
	h.Write([]byte(data))

	return h.Sum(nil)
}
