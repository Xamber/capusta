package capusta

import (
	"sync"
)

type Wallet struct {
	owner        string
	blockchain   *blockchain
	transactions map[[32]byte]float64
	lock         sync.Mutex
}

func (w *Wallet) GetBalance() float64 {
	allMoney := 0.0
	w.RefreshTransactions()

	w.lock.Lock()
	defer w.lock.Unlock()

	for _, money := range w.transactions {
		allMoney += money
	}
	return allMoney
}

func (w *Wallet) TransferMoney(to *Wallet, amount float64) (string, error) {

	w.RefreshTransactions()

	w.lock.Lock()
	defer w.lock.Unlock()

	transfered := 0.00

	inputs := []TInput{}
	outputs := []TOutput{}

	for id, money := range w.transactions {
		inputs = append(inputs, TInput{id, money, w.owner})
		transfered += money
	}

	if transfered < amount {
		return "", ErrorNotEnoghtMoney
	}

	outputs = append(outputs, TOutput{value: amount, to: to.owner})

	change := transfered - amount
	if change != 0.0 {
		outputs = append(outputs, TOutput{value: change, to: w.owner})
	}

	transaction := NewTransaction(inputs, outputs)
	w.blockchain.transactions = append(w.blockchain.transactions, transaction)

	return transaction.getID(), nil
}

func (w *Wallet) CheckTransactionOwner(t Transaction) (bool, float64, bool) {
	var owner bool = false
	var haveOutput float64 = 0
	var haveInput bool = false

	for _, out := range t.outputs {
		if out.to == w.owner {
			owner = true
			haveOutput += out.value
		}
	}

	for _, in := range t.inputs {
		if in.from == w.owner {
			owner = true
			haveInput = true
		}
	}
	return owner, haveOutput, haveInput
}

func (w *Wallet) RefreshTransactions() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.transactions = map[[32]byte]float64{}
	for curBlock := range w.blockchain.Iterator() {
		for _, tr := range curBlock.GetTransactions() {

			owner, money, used := w.CheckTransactionOwner(tr)
			if owner == false {
				continue
			}

			if money > 0 {
				w.transactions[tr.hash] = w.transactions[tr.hash] + money
			}
			if _, ok := w.transactions[tr.hash]; ok && used {
				delete(w.transactions, tr.hash)
			}
		}
	}
}
