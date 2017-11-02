package capusta


type Transaction struct {
	sender   string
	receiver string
	amount   float64
}

type transactions []Transaction

func NewTransaction(sender string, receiver string, amount float64) Transaction {
	t := Transaction{sender, receiver, amount}
	UpcomingTransaction = append(UpcomingTransaction, t)
	return t
}
