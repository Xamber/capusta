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
		if block.hash == hash{
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
	chain.transactions = append(chain.transactions, createRewardTransaction(miner))

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
	ownerTransactions := map[string]Transaction{}

	for index := 1; index < chain.getLenght(); index++ {
		currentBlock := chain.getBlockbyIndex(index)

		for _, transaction := range currentBlock.GetTransactions() {

			for _, out := range transaction.Outputs {
				if out.To == owner {
					ownerTransactions[transaction.getID()] = transaction
					break
				}
			}

			for _, in := range transaction.Inputs {
				if _, ok := ownerTransactions[in.TransactionID]; ok && in.From == owner {
					delete(ownerTransactions, in.TransactionID)
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
func (chain *blockchain) TransferMoney(from, to string, amount float64) (string, error) {

	preperadTransactions := map[string]float64{}
	money := 0.0000

	for _, t := range chain.FindAvalibleTransactions(from) {
		for _, o := range t.Outputs {
			if !(o.To == from) {
				continue
			}

			money = money + o.Value
			preperadTransactions[t.getID()] = o.Value
		}

		if money >= amount {
			break
		}
	}

	if money < amount {
		return "", ErrorNotEnoghtMoney
	}

	inputs := []Input{}
	outputs := []Output{}

	for id, value := range preperadTransactions {
		inputs = append(inputs, Input{id, value, from})
	}

	outputs = append(outputs, Output{amount, to})

	if money != amount {
		outputs = append(outputs, Output{money - amount, from})
	}

	transaction := Transaction{defaultHash32, inputs, outputs}
	transaction.setHandlers()
	chain.transactions = append(chain.transactions, transaction)

	return transaction.getID(), nil

}

// impliment Stringer interface
func (chain blockchain) String() string {

	var blocksInfo string = ""

	for b := range chain.Iterator() {
		blocksInfo += b.String()
	}

	return fmt.Sprintf("Blockchain - Length: %d \n%s", chain.getLenght(), blocksInfo)
}
