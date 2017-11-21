package capusta

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
	"errors"
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

func (chain *blockchain) FindAvalibleTransactions(owner string) []transaction {
	ownerTransactions := map[string]transaction{}

	for index := 1; index < chain.getLenght(); index++ {
		currentBlock := chain.getBlockbyIndex(index)

		for _, transaction := range currentBlock.getTransactions() {

			for _, out := range transaction.outputs {
				if out.validateOwner(owner) {
					ownerTransactions[transaction.id] = transaction
					break
				}
			}

			for _, in := range transaction.inputs {
				if _, ok := ownerTransactions[in.transactionID]; ok && in.validateOwner(owner) {
					delete(ownerTransactions, in.transactionID)
				}
			}
		}
	}

	unspendTransaction := []transaction{}
	for _, v := range ownerTransactions {
		unspendTransaction = append(unspendTransaction, v)
	}

	return unspendTransaction
}

func (chain *blockchain) TransferMoney(from, to string, amount float64) (string, error) {

	preperadTransactions := map[string]float64{}
	money := 0.0000

	for _, t := range chain.FindAvalibleTransactions(from) {
		for _, o := range t.outputs {

			if !o.validateOwner(from) {
				continue
			}

			money += o.value
			preperadTransactions[t.id] = o.value
		}

		if money >= amount {
			break
		}
	}

	if money < amount {
		return "", errors.New("User don't have enough money")
	}

	inputs := []input{}
	outputs := []output{}

	for id, value := range preperadTransactions {
		inputs = append(inputs, input{id, value, from})
	}

	outputs = append(outputs, output{amount, to})

	if money != amount {
		outputs = append(outputs, output{money - amount, from})
	}

	transaction := transaction{"", defaultHash, inputs, outputs}
	transaction.setHandlers()
	chain.transactions = append(chain.transactions, transaction)

	return transaction.id, nil

}

func (chain *blockchain) Info() {

	fmt.Printf("Blockchain - Length: %v \n", chain.getLenght())

	for i := chain.getLenght(); i > 1; i-- {
		fmt.Println(chain.blocks[i-1].info())
	}

}
