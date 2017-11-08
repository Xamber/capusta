package capusta

import (
	"encoding/json"
)

// transaction impliment simple transaction entity
type transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

// transactions is a list of transactions
type transactions []transaction

// serialize create bytes from structure
func (t *transactions) serialize() string {
	seriliazed, _ := json.Marshal(t)
	return string(seriliazed)
}
