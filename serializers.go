package capusta

import (
	"bytes"
	"encoding/gob"
)

// SerializeTransactions create bytes list of transactions
func SerializeTransactions(ts []Transaction) []byte {
	result :=  bytes.Buffer{}
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(ts)
	handleError(err)

	return result.Bytes()
}

// DeserializeTransactions deserializes a list of transactions
func DeserializeTransactions(binary []byte) []Transaction {
	ts := []Transaction{}

	decoder := gob.NewDecoder(bytes.NewReader(binary))
	err := decoder.Decode(ts)
	handleError(err)

	return ts
}
