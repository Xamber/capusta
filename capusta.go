package capusta

import (
	"time"
)

var DEFAULT_PROOF = []byte{0, 0}

var Blockchain blockchain

func init() {

	genesisBlock := Block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         "",
		proof:        1337,
		previousHash: [32]byte{},
		hash:         [32]byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}
