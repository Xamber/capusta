package capusta

import (
	"encoding/json"
)

type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type Transactions []Transaction

func (t *Transactions) Serialize() string {
	seriliazed, _ := json.Marshal(t)
	return string(seriliazed)
}
