package capusta

import (
	"time"
	"sync"
	"fmt"
	"encoding/json"
	"bytes"
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

func (b *blockchain) AddBlock(proof int) {

	data := b.transactions.Serialize()

	block := Block{
		index:        len(b.blocks),
		timestamp:    time.Now().UnixNano(),
		data:         data,
		proof:        proof,
		previousHash: Hash(b.getLastBlock().data),
	}

	b.blocks = append(b.blocks, block)
	b.transactions = Transactions{}
}

func (b *blockchain) MineBlock() {

	defaultProof := []byte(DEFAULT_PROOF)
	defaultProofLenght := len(defaultProof)
	lastProof := b.getLastBlock().proof
	proof := 1

	for {
		hashString := fmt.Sprintf("%d%d", lastProof, proof)
		hash := Hash(hashString)[:defaultProofLenght]

		if bytes.Equal(hash, defaultProof) {
			break
		}

		proof += 1
	}

	b.AddBlock(proof)
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
		fmt.Println(v.index, v.timestamp, "previous hash:", v.previousHash, "currient hash:", Hash(v.data))
		fmt.Println(v.data)
		fmt.Println()
	}

}
