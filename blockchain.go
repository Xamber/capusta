package capusta

import (
	"fmt"
	"sync"
	"time"
)

type blockchain struct {
	blocks       []Block
	transactions []Transaction
	lock         sync.Mutex
}

// blockchain.getLenght return lenght of blockchain
func (chain *blockchain) getLenght() int {
	return len(chain.blocks)
}

// blockchain.getLastBlock return pointer to last Block of blockchain
func (chain *blockchain) getLastBlock() *Block {
	return &chain.blocks[chain.getLenght()-1]
}

// blockchain.getBlockbyIndex return pointer to Block by index (int)
func (chain *blockchain) getBlockbyIndex(index int) *Block {
	return &chain.blocks[index]
}

// blockchain.getBlockbyHash search and return pointer to Block by hash ([32]byte)
func (chain *blockchain) getBlockbyHash(hash [32]byte) *Block {
	for block := range chain.Iterator() {
		if block.hash == hash {
			return block
		}
	}

	return nil
}

func (chain *blockchain) GetWallet(owner string) *Wallet {
	return &Wallet{owner: owner, blockchain: chain}
}

func (chain *blockchain) Iterator() chan *Block {
	output := make(chan *Block)

	push := func(block Block) {
		output <- &block
	}

	pusher := func() {
		for _, v := range chain.blocks {
			push(v)
		}
		close(output)
	}

	go pusher()
	return output
}

// blockchain.MineBlock mine Block
// Function lock transaction, create reward for miner and starting for searching proof-of-work
func (chain *blockchain) MineBlock(miner string) {
	chain.transactions = append(chain.transactions, NewReward(miner))

	var block = Block{
		index:        int64(chain.getLenght()),
		timestamp:    time.Now().UnixNano(),
		data:         chain.transactions,
		previousHash: chain.getLastBlock().hash,
		proof:        1,
	}

	for {
		block.proof++
		block.hash = Hash(&block)

		if block.checkSum() {
			break
		}
	}

	chain.transactions = []Transaction{}
	chain.blocks = append(chain.blocks, block)
}

// impliment Stringer interface
func (chain blockchain) String() string {

	var blocksInfo string = ""

	for b := range chain.Iterator() {
		blocksInfo += b.String()
	}

	return fmt.Sprintf("Blockchain - Length: %d \n%s", chain.getLenght(), blocksInfo)
}
