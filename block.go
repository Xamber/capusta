package capusta

import (
	"bytes"
	"fmt"
)

// Block contain information about Block
type Block struct {
	index        int64
	timestamp    int64
	data         []Transaction
	proof        int64
	hash         [32]byte
	previousHash [32]byte
}

func (b *Block) GetTransactions() []Transaction {
	return b.data
}

// Block.validate check hash of Block
func (b *Block) validate() bool {
	hash := Hash(b)
	return bytes.Equal(hash[:], b.hash[:])
}

func (b *Block) checkSum() bool {
	return bytes.HasPrefix(b.hash[:], defaultProof)
}

// Block.info return string with info about Block
func (b *Block) String() string {

	transactions := ""
	for _, t := range b.data {
		transactions += t.String()
	}

	if transactions == "" {
		transactions = "<Empty transactions>"
	}

	template := "Block %v \nTimestamp: %v Proof: %v \nHash: %x\nPreviousHash: %x\nValidated: %v\nTransactions: %v\n\n"
	return fmt.Sprintf(template, b.index, b.timestamp, b.proof, b.hash, b.previousHash, b.validate(), transactions)
}
