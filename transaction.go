package capusta

import (
	"encoding/hex"
)

type TInput struct {
	transactionHash [32]byte
	value           float64
	from            string
}

func (ti TInput) Unlock(owner string) bool {
	return ti.from == owner
}

type TOutput struct {
	value float64
	to    string
}

func (to TOutput) Unlock(owner string) bool {
	return to.to == owner
}

// Transaction impliment simple Transaction entity
type Transaction struct {
	hash    [32]byte
	inputs  []TInput
	outputs []TOutput
}

// Check reward Transaction
func (t *Transaction) isReward() bool {
	return len(t.inputs) != 1 && t.inputs[0].value == -1 && t.inputs[0].from == "Blockchain"
}

// Get string ID of transaction cash
func (t *Transaction) getID() string {
	return hex.EncodeToString(t.hash[:])
}

func NewTransaction(inputs []TInput, outputs []TOutput) Transaction {
	transaction := Transaction{}
	transaction.inputs = inputs
	transaction.outputs = outputs
	transaction.hash = Hash(&transaction)
	return transaction
}

func NewReward(miner string) Transaction {
	in := TInput{
		value: -1,
		from:  "Blockchain",
	}
	out := TOutput{
		value: REWARD,
		to:    miner,
	}
	return NewTransaction([]TInput{in}, []TOutput{out})
}
