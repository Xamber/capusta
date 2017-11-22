package capusta

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"encoding/binary"
)

// block contain information about block
type block struct {
	index        int
	timestamp    int64
	data         []Transaction
	proof        int64
	hash         [32]byte
	previousHash [32]byte
}

func (b *block) getTransactions() []Transaction {
	return b.data
}

func (b *block) makeBLOB(proof int64) []byte {
	var binaryData bytes.Buffer
	var serializedTransaction []byte

	serializedTransaction = SerializeTransactions(b.getTransactions())

	write := func(add interface{}) {
		err := binary.Write(&binaryData, binary.LittleEndian, add)
		handleError(err)
	}

	write(b.previousHash)
	write(b.timestamp)
	write(serializedTransaction)
	write(proof)

	return binaryData.Bytes()
}

func (b *block) makeHash(proof int64) [32]byte {
	return sha256.Sum256(b.makeBLOB(proof))
}

// block.validate check Hash of block
func (b *block) validate() bool {
	hash := b.makeHash(b.proof)
	return bytes.HasPrefix(hash[:], b.hash[:])
}

// block.info return string with info about block
func (b *block) String() string {
	template := "block Index: %v Timestamp: %v Proof: %v \nHash: %x\nPreviousHash: %x\nValidated: %v\nTransactions: %v\n\n"
	return fmt.Sprintf(template, b.index, b.timestamp, b.proof, b.hash, b.previousHash, b.validate(), b.data)
}
