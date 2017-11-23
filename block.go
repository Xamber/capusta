package capusta

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// Block contain information about Block
type Block struct {
	index        int
	timestamp    int64
	data         []Transaction
	proof        int64
	hash         [32]byte
	previousHash [32]byte
}

func (b *Block) GetTransactions() []Transaction {
	return b.data
}

func (b *Block) toBinary(proof int64) []byte {
	var binaryData bytes.Buffer
	var serializedTransaction []byte

	serializedTransaction = SerializeTransactions(b.GetTransactions())

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

func (b *Block) Hash(proof int64) [32]byte {
	return sha256.Sum256(b.toBinary(proof))
}

func (b *Block) transactionToBinary() []byte {
	blob := [][]byte{}

	for _, t := range b.data {
		blob = append(blob, t.makeBLOB())
	}

	return bytes.Join(blob, []byte{})
}

func (b *Block) hashTransactions() [32]byte {
	return sha256.Sum256(b.transactionToBinary())
}

// Block.validate check Hash of Block
func (b *Block) validate() bool {
	hash := b.Hash(b.proof)
	return bytes.HasPrefix(hash[:], b.hash[:])
}

// Block.info return string with info about Block
func (b *Block) String() string {
	template := "Block %v \nTimestamp: %v Proof: %v \nHash: %x\nPreviousHash: %x\nValidated: %v\nTransactions: %v\n\n"
	return fmt.Sprintf(template, b.index, b.timestamp, b.proof, b.hash, b.previousHash, b.validate(), b.data)
}
