package capusta

type Block struct {
	index        int
	timestamp    int64
	data         string
	proof        int
	previousHash []byte
}
