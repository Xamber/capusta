package capusta

import (
	"time"
)

var defaultProof = []byte{0, 0}

// Blockchain is the global blockchain variable
var Blockchain blockchain

func init() {

	genesisBlock := block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         "",
		proof:        1337,
		previousHash: [32]byte{},
		hash:         [32]byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}
