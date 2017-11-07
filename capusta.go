package capusta

import (
	"time"
)

var Blockchain blockchain

func init() {

	genesisBlock := Block{
		index:         0,
		timestamp:     time.Now().Unix(),
		data:          "",
		proof:         100,
		previousBlock: nil,
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}
