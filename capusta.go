package capusta

import (
	"time"
	"errors"
)

var defaultProof = []byte{0, 0}
var defaultHash32 = [32]byte{}
var defaultHash = defaultHash32[:]

const REWARD = 1000

var ErrorNotEnoghtMoney = errors.New("User don't have enough money")

// Blockchain is the global blockchain variable
var Blockchain blockchain

func init() {

	genesisBlock := block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         []Transaction{},
		proof:        1337,
		previousHash: [32]byte{},
		hash:         [32]byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}
