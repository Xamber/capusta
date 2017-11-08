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

func (b *blockchain) AddBlock(block Block) {
	b.blocks = append(b.blocks, block)
}

func (b *blockchain) MineBlock() {

	block := Block{
		index:        len(b.blocks),
		timestamp:    time.Now().UnixNano(),
		data:         b.transactions.Serialize(),
		previousHash: b.getLastBlock().hash,
	}

	defaultProofLenght := len(DEFAULT_PROOF)

	var proof int64 = 1
	var hash = [32]byte{}

	for {
		data := block.PrepareData(proof)
		hash = sha256.Sum256(data)

		if bytes.Equal(hash[:defaultProofLenght], DEFAULT_PROOF) {
			break
		}

		proof += 1
	}

	block.proof = proof
	block.hash = hash

	b.AddBlock(block)
}

func (b *blockchain) AddTransaction(sender string, receiver string, amount float64) Transaction {
	t := Transaction{sender, receiver, amount}
	b.transactions = append(b.transactions, t)
	return t
}

func (b *blockchain) getLastBlock() *Block {
	lastBlock := b.blocks[len(b.blocks)-1]
	return &lastBlock
}

func (b *blockchain) Log() {

	for i := len(b.blocks); i > 1; i-- {
		v := b.blocks[i-1]
		fmt.Println(v.index, v.timestamp)
		fmt.Printf("Previous Hash: %x\n", v.previousHash)
		fmt.Printf("Hash: %x\n", v.hash)
		fmt.Println(v.data)
		fmt.Println(v.ValidateHash())
		fmt.Println()
	}

}
