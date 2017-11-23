package capusta

import (
	"errors"
	"time"
	"log"
)

var defaultProof = []byte{0, 0}
var defaultHash32 = [32]byte{}
var defaultHash = defaultHash32[:]

const REWARD = 1000

var ErrorNotEnoghtMoney = errors.New("User don't have enough money")

var Blockchain blockchain

func init() {

	genesisBlock := Block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         []Transaction{},
		proof:        1337,
		previousHash: [32]byte{},
		hash:         [32]byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}

func handleError(err error)  {
	if err != nil {
		log.Panic(err)
	}
}
