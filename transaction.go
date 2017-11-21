package capusta

import (
	"bytes"
	"encoding/gob"
)

type input struct {
	transactionHash [32]byte
	value           int
	from            string
}

type output struct {
	value int
	to    string
}

// transaction impliment simple transaction entity
type transaction struct {
	hash    [32]byte
	inputs  []input
	outputs []output
}

// transactions is a list of transactions
type transactions []transaction

func createRewardTransaction(miner string) transaction {
	in := input{defaultHash, -1, "Blockchain"}
	out := output{REWARD, miner}
	transaction := transaction{nil, []input{in}, []output{out}}
	transaction.setHash()
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
func (t *transaction) setHash() {
	t.hash = Hashing(t)
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
