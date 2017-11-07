package capusta

import (
	"time"
	"sync"
	"fmt"
	"encoding/json"
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

func (b *blockchain) AddBlock() {

	block := Block{
		index:         len(b.blocks),
		timestamp:     time.Now().UnixNano(),
		data:          b.transactions.Serialize(),
		proof:         3,
		previousBlock: b.getLastBlock(),
	}
	b.blocks = append(b.blocks, block)
	b.transactions = Transactions{}
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
		fmt.Println(v.index, v.timestamp, "previous hash:", v.previousBlock.hash())
	}

}
