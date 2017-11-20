package capusta

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

type blockchain struct {
	blocks       []block
	transactions transactions
	lock         sync.Mutex
}

func (chain *blockchain) getLenght() int {
	return len(chain.blocks)
}

func (chain *blockchain) getLastBlock() *block {
	return &chain.blocks[chain.getLenght()-1]
}

func (chain *blockchain) getBlockbyIndex(index int) *block {
	return &chain.blocks[index]
}

func (chain *blockchain) getBlockbyHash(hash [32]byte) *block {
	for i := chain.getLenght() - 1; i >= 0; i-- {
		block := chain.getBlockbyIndex(i)
		if block.hash == hash {
			return block
		}
	}
	return nil
}

func (chain *blockchain) MineBlock(miner string) {

	var proof int64 = 1
	var hash = [32]byte{}

	chain.transactions = append(chain.transactions, createRewardTransaction(miner))

	var block = block{
		index:        chain.getLenght(),
		timestamp:    time.Now().UnixNano(),
		data:         chain.transactions.serialize(),
		previousHash: chain.getLastBlock().hash,
	}

	for {
		data := block.prepareData(proof)
		hash = sha256.Sum256(data)

		if bytes.HasPrefix(hash[:], defaultProof) {
			break
		}

		proof++
	}

	block.proof = proof
	block.hash = hash

	chain.transactions = transactions{}
	chain.blocks = append(chain.blocks, block)
}

func (chain *blockchain) AddTransaction(sender string, receiver string, amount float64) transaction {
	t := transaction{sender, receiver, amount}
	chain.transactions = append(chain.transactions, t)
	return t
}

func (chain *blockchain) Info() {

	fmt.Printf("Blockchain - Length: %v \n", chain.getLenght())

	for i := chain.getLenght(); i > 1; i-- {
		fmt.Println(chain.blocks[i-1].info())
	}

}
