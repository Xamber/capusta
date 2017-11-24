package capusta

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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

func (ti TInput) DataToBinary() []any {
	return []any{ti.TransactionHash, ti.Value, []byte(ti.From)}
}

type TOutput struct {
	Value float64
	To    string
}

func (to TOutput) DataToBinary() []any {
	return []any{to.Value, []byte(to.To)}
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

func (t *Transaction) DataToBinary() []any {
	data := []any{}

	for _, i := range t.Inputs {
		data = append(data, Binarizate(i))
	}

	for _, o := range t.Outputs {
		data = append(data, Binarizate(o))
	}

	return data
}

func (t *Transaction) makeBLOB() []byte {
	var binaryData bytes.Buffer

	encoder := gob.NewEncoder(&binaryData)

	err := encoder.Encode(*t)
	handleError(err)

	return binaryData.Bytes()
}

func (t *Transaction) makeHash() [32]byte {
	return sha256.Sum256(t.makeBLOB())
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
