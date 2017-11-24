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
		index:        chain.getLenght(),
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

// blockchain.FindAvalibleTransactions find avalible transaction for create another one
func (chain *blockchain) FindAvalibleTransactions(owner string) []Transaction {
	ownerTransactions := map[[32]byte]Transaction{}

	for currentBlock := range chain.Iterator() {

		for _, transaction := range currentBlock.GetTransactions() {

			for _, out := range transaction.Outputs {
				if out.Unlock(owner) {
					ownerTransactions[transaction.Hash] = transaction
					break
				}
			}

			for _, in := range transaction.Inputs {
				if _, ok := ownerTransactions[in.TransactionHash]; ok && in.Unlock(owner) {
					delete(ownerTransactions, in.TransactionHash)
				}
			}
		}
	}

	unspendTransaction := []Transaction{}
	for _, v := range ownerTransactions {
		unspendTransaction = append(unspendTransaction, v)
	}

	return unspendTransaction
}

// blockchain.TransferMoney create transaction in blockcahin
func (chain *blockchain) TransferMoney(from, to string, amount float64) ([32]byte, error) {

	preperadTransactions := map[[32]byte]float64{}
	money := 0.0000

	for _, t := range chain.FindAvalibleTransactions(from) {
		for _, o := range t.Outputs {
			if !(o.To == from) {
				continue
			}

			money = money + o.Value
			preperadTransactions[t.Hash] = o.Value
		}

		if money >= amount {
			break
		}
	}

	if money < amount {
		return [32]byte{}, ErrorNotEnoghtMoney
	}

	inputs := []TInput{}
	outputs := []TOutput{}

	for id, value := range preperadTransactions {
		inputs = append(inputs, TInput{id, value, from})
	}

	outputs = append(outputs, TOutput{amount, to})

	if money != amount {
		outputs = append(outputs, TOutput{money - amount, from})
	}
	transaction := NewTransaction(inputs, outputs)
	chain.transactions = append(chain.transactions, transaction)

	return transaction.Hash, nil

}

// impliment Stringer interface
func (chain blockchain) String() string {

	var blocksInfo string = ""

	for b := range chain.Iterator() {
		blocksInfo += b.String()
	}

	return fmt.Sprintf("Blockchain - Length: %d \n%s", chain.getLenght(), blocksInfo)
}
