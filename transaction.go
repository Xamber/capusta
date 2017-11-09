package capusta

import (
	"bytes"
	"encoding/gob"
)

// transaction impliment simple transaction entity
type transaction struct {
	Sender   string
	Receiver string
	Amount   float64
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

//  DeserializeBlock deserializes a block
func (t *transactions) deserialize(binary []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(binary))
	err := decoder.Decode(t)
	logError(err)
}
