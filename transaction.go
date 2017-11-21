package capusta

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

type Input struct {
	TransactionID string
	Value         float64
	From          string
}

func (i *Input) validateOwner(owner string) bool {
	return i.From == owner
}

type Output struct {
	Value float64
	To    string
}

func (o *Output) validateOwner(owner string) bool {
	return o.To == owner
}

// Transaction impliment simple Transaction entity
type Transaction struct {
	ID      string
	Hash    [32]byte
	Inputs  []Input
	Outputs []Output
}

// transactions is a list of transactions
type transactions []Transaction

func createRewardTransaction(miner string) Transaction {
	in := Input{"", -1, "Blockchain"}
	out := Output{REWARD, miner}
	transaction := Transaction{"", defaultHash, []Input{in}, []Output{out},}
	transaction.setHandlers()
	return transaction
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
	t.Hash = Hashing(t)
	t.ID = hex.EncodeToString(t.Hash[:])
}

// serialize create bytes From structure
func (ts *transactions) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(ts)
	logError(err)

	return result.Bytes()
}

// deserialize deserializes a list of transactions
func (ts *transactions) deserialize(binary []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(binary))
	err := decoder.Decode(ts)
	logError(err)
}
