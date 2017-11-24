package capusta

import (
	"errors"
	"log"
	"time"
)

var Blockchain blockchain

var defaultProof = []byte{0, 0}
var defaultHash32 = [32]byte{}
var defaultHash = defaultHash32[:]

const REWARD = 1000

type any interface{}

var ErrorNotEnoghtMoney = errors.New("User don't have enough money")

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

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
