package capusta

import (
	"time"
	"crypto/sha256"
	"bytes"
	"encoding/binary"
)

var DEFAULT_PROOF = []byte{0, 0, 0}

var Blockchain blockchain

func init() {

	genesisBlock := Block{
		index:        0,
		timestamp:    time.Now().UnixNano(),
		data:         "",
		proof:        1337,
		previousHash: [32]byte{},
	}

	Blockchain.blocks = append(Blockchain.blocks, genesisBlock)
}

func Hash(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func toBinary(in int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, in)
	return buf.Bytes()
}
