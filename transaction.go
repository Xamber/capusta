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

// serialize create bytes from structure
func (t *transactions) serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(t)
	logError(err)

	return result.Bytes()
}

// deserialize deserializes a list of transactions
func (t *transactions) deserialize(binary []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(binary))
	err := decoder.Decode(t)
	logError(err)
}
