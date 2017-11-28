package capusta

import (
	"encoding/hex"
	"fmt"
)

type TInput struct {
	transactionHash [32]byte
	value           float64
	from            string
}

type TOutput struct {
	value float64
	to    string
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
	transaction := Transaction{inputs: inputs, outputs: outputs}
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

func (t *Transaction) String() string {
	var ret string = fmt.Sprintf("\n    %s\n", t.getID())

	for _, ti := range t.inputs {
		ret += fmt.Sprintf("    IN %s - %f\n", ti.from, ti.value)
	}

	for _, to := range t.outputs {
		ret += fmt.Sprintf("    OUT %s - %f\n", to.to, to.value)
	}

	return ret
}
