package capusta

import (
	"time"
	"sync"
	"fmt"
	"encoding/json"
	"bytes"
	"crypto/sha256"
)

type blockchain struct {
	blocks       []Block
	transactions Transactions
	lock         sync.Mutex
}

type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type Transactions []Transaction

func (t *Transactions) Serialize() string {
	seriliazed, _ := json.Marshal(t)
	return string(seriliazed)
}

func (chain *blockchain) getLenght() int {
	return len(chain.blocks)
}

func (chain *blockchain) getLastBlock() *Block {
	return &chain.blocks[chain.getLenght()-1]
}

func (chain *blockchain) MineBlock() {

	var proof int64 = 1
	var hash = [32]byte{}

	var block = Block{
		index:        len(chain.blocks),
		timestamp:    time.Now().UnixNano(),
		data:         chain.transactions.Serialize(),
		previousHash: chain.getLastBlock().hash,
	}

	for {
		data := block.PrepareData(proof)
		hash = sha256.Sum256(data)

		if bytes.Equal(hash[:len(DEFAULT_PROOF)], DEFAULT_PROOF) {
			break
		}

		proof += 1
	}

	block.proof = proof
	block.hash = hash

	chain.transactions = Transactions{}
	chain.blocks = append(chain.blocks, block)
}

func (chain *blockchain) AddTransaction(sender string, receiver string, amount float64) Transaction {
	t := Transaction{sender, receiver, amount}
	chain.transactions = append(chain.transactions, t)
	return t
}

func (chain *blockchain) Log() {

	for i := chain.getLenght(); i > 1; i-- {
		fmt.Println(chain.blocks[i-1].Info())
	}

}
