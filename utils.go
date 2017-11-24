package capusta

import (
	"bytes"
	"encoding/binary"
	"crypto/sha256"
)

type Hashible interface {
	DataToBinary() []any
}

func Binarizate(obj Hashible) []byte {
	var binaryData bytes.Buffer

	write := func(add any) {
		err := binary.Write(&binaryData, binary.LittleEndian, add)
		handleError(err)
	}

	for _, data := range obj.DataToBinary() {
		write(data)
	}

	return binaryData.Bytes()
}

func Hash(obj Hashible) [32]byte {
	data := Binarizate(obj)
	hash := sha256.Sum256(data)
	return hash
}
