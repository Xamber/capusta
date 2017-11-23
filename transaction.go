package capusta

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

type Input struct {
	TransactionID string
	Value         float64
	From          string
}

type Output struct {
	Value float64
	To    string
}

// Transaction impliment simple Transaction entity
type Transaction struct {
	Hash    [32]byte
	Inputs  []Input
	Outputs []Output
}

// Check reward Transaction
func (t *Transaction) isReward() bool {
	if len(t.Inputs) != 1 {
		return false
	}
	in := t.Inputs[0]
	return in.Value == -1 && in.From == "Blockchain"
}

// Set Hash To Transaction
func (t *Transaction) setHandlers() {
	t.Hash = t.makeHash()
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

func createTransaction(inputs []Input, outputs []Output) Transaction {
	transaction := Transaction{}
	transaction.Inputs = inputs
	transaction.Outputs = outputs
	transaction.setHandlers()
	return transaction
}

func createRewardTransaction(miner string) Transaction {
	in := Input{"", -1, "Blockchain"}
	out := Output{REWARD, miner}
	return createTransaction([]Input{in}, []Output{out})
}
