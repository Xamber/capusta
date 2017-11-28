package capusta

import (
	"errors"
	"log"
	"time"
)

var Blockchain blockchain

var defaultProof = []byte{0, 0}

const REWARD = 1000

var ErrorNotEnoghtMoney = errors.New("User don't have enough money")
var ErrorWalletDontHaveBlockchain = errors.New("Wallet don't have connection with Blockchain")

func init() {

	genesisBlock := Block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         []Transaction{},
		proof:        1337,
		previousHash: [32]byte{},
	}

	genesisBlock.hash = Hash(&genesisBlock)
	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
