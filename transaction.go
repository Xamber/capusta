package capusta

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

type input struct {
	transactionID 	string
	value           float64
	from            string
}

func (i *input) validateOwner(owner string) bool {
	return i.from == owner
}

type output struct {
	value float64
	to    string
}

func (o *output) validateOwner(owner string) bool {
	return o.to == owner
}

// transaction impliment simple transaction entity
type transaction struct {
	id      string
	hash    [32]byte
	inputs  []input
	outputs []output
}

// transactions is a list of transactions
type transactions []transaction

func createRewardTransaction(miner string) transaction {
	in := input{"", -1, "Blockchain"}
	out := output{REWARD, miner}
	transaction := transaction{"", defaultHash, []input{in}, []output{out},}
	transaction.setHandlers()
	return transaction
}

// Check reward transaction
func (t *transaction) isReward() bool {
	if len(t.inputs) != 1 {
		return false
	}
	in := t.inputs[0]
	return in.value == -1 && in.from == "Blockchain"
}

// Set hash to transaction
func (t *transaction) setHandlers() {
	t.hash = Hashing(t)
	t.id = hex.EncodeToString(t.hash[:])
}

// serialize create bytes from structure
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
