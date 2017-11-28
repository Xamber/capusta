package capusta

import (
	"encoding/hex"
)

type TInput struct {
	TransactionHash [32]byte
	Value           float64
	From            string
}

func (ti TInput) Unlock(owner string) bool {
	return ti.From == owner
}

type TOutput struct {
	Value float64
	To    string
}

func (to TOutput) Unlock(owner string) bool {
	return to.To == owner
}

// Transaction impliment simple Transaction entity
type Transaction struct {
	Hash    [32]byte
	Inputs  []TInput
	Outputs []TOutput
}

// Check reward Transaction
func (t *Transaction) isReward() bool {
	return len(t.Inputs) != 1 && t.Inputs[0].Value == -1 && t.Inputs[0].From == "Blockchain"
}

// Get string ID of transaction cash
func (t *Transaction) getID() string {
	return hex.EncodeToString(t.Hash[:])
}

func NewTransaction(inputs []TInput, outputs []TOutput) Transaction {
	transaction := Transaction{}
	transaction.Inputs = inputs
	transaction.Outputs = outputs
	transaction.Hash = Hash(&transaction)
	return transaction
}

func NewReward(miner string) Transaction {
	in := TInput{
		Value: -1,
		From:  "Blockchain",
	}
	out := TOutput{
		Value: REWARD,
		To:    miner,
	}
	return NewTransaction([]TInput{in}, []TOutput{out})
}
