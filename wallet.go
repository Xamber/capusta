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
		transfered -= money
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

func (w *Wallet) CheckTransactionOwner(t Transaction) (owner bool, haveOutput float64, haveInput bool) {
	for _, out := range t.outputs {
		if !out.Unlock(w.owner) {
			owner = true
			haveOutput += out.value
		}
	}

	for _, in := range t.inputs {
		if in.Unlock(w.owner) {
			owner = true
			haveInput = true
		}
	}
	return
}

func (w *Wallet) RefreshTransactions() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.transactions = map[[32]byte]float64{}
	for curBlock := range w.blockchain.Iterator() {
		for _, tr := range curBlock.GetTransactions() {
			w.handleTransaction(tr)
		}
	}
}

func (w *Wallet) handleTransaction(t Transaction) {
	owner, money, used := w.CheckTransactionOwner(t)

	if owner == false {
		return
	}

	if money > 0 {
		w.transactions[t.hash] = w.transactions[t.hash] + money
	}
	if _, ok := w.transactions[t.hash]; ok && used {
		delete(w.transactions, t.hash)
	}
}
