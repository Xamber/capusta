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

type TOutput struct {
	Value float64
	To    string
}

// Transaction impliment simple Transaction entity
type Transaction struct {
	Hash    [32]byte
	Inputs  []TInput
	Outputs []TOutput
}

// Check reward Transaction
func (t *Transaction) isReward() bool {
	if len(t.Inputs) != 1 {
		return false
	}
	in := t.Inputs[0]
	return in.Value == -1 && in.From == "Blockchain"
}

// Get string ID of transaction cash
func (t *Transaction) getID() string {
	return hex.EncodeToString(t.Hash[:])
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
	transaction.Hash = transaction.makeHash()
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
