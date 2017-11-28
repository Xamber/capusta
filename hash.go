package capusta

import (
	"crypto/sha256"
	"encoding/binary"
	"math"
)

type Hashible interface {
	Binary() []byte
}

func Hash(obj Hashible) [32]byte {
	data := obj.Binary()
	return sha256.Sum256(data)
}

type blob struct {
	bytes []byte
}

func (b *blob) Bytes() []byte {
	return b.bytes
}

func (b *blob) Write(chunk []byte) {
	b.bytes = append(b.bytes, chunk...)
}

func (b *blob) WriteHash(chunk [32]byte) {
	b.Write(chunk[:])
}

func (b *blob) WriteString(chunk string) {
	b.Write([]byte(chunk))
}

func (b *blob) WriteInt64(chunk int64) {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(chunk))
	b.Write(bytes)
}

func (b *blob) WriteFloat64(chunk float64) {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, math.Float64bits(chunk))
	b.Write(bytes)
}

func (b *Block) Binary() []byte {

	var data blob
	var transactionData blob

	for _, t := range b.data {
		transactionData.Write(t.Binary())
	}

	data.WriteInt64(b.index)
	data.WriteInt64(b.timestamp)
	data.WriteHash(b.previousHash)
	data.Write(transactionData.bytes)
	data.WriteInt64(b.proof)

	return data.Bytes()
}

func (t *Transaction) Binary() []byte {
	var data blob

	for _, ti := range t.Inputs {
		data.WriteHash(ti.TransactionHash)
		data.WriteFloat64(ti.Value)
		data.WriteString(ti.From)
	}

	for _, to := range t.Outputs {
		data.WriteFloat64(to.Value)
		data.WriteString(to.To)
	}

	return data.Bytes()
}
