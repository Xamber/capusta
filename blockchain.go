package capusta

import (
	"time"
	"sync"
)

type blockchain struct {
	blocks []Block
	lock	sync.Mutex
}

func (b *blockchain) AddBlock() {

	block := Block{
		index:         len(*b) + 1,
		timestamp:     time.Now().Unix(),
		data:          UpcomingTransaction,
		proof:         3,
		previousBlock: b.getLastBlock(),
	}

	*b = append((*b), block)
}

func (b *blockchain) getLastBlock() *Block {

	blockchainLeight := len(*b)
	lastBlock := &(*b)[blockchainLeight+1]

	return lastBlock
}
