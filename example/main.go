package main

import (
	"github.com/xamber/capusta"
)

var blockchain = capusta.Blockchain

func main() {

	blockchain.AddTransaction("Artem", "Dima", 100)
	blockchain.AddTransaction("Dima", "Artem", 50)
	blockchain.AddTransaction("Artem", "Dima", 80)

	blockchain.AddBlock()

	blockchain.AddTransaction("Dima", "Artem", 77)
	blockchain.AddTransaction("Artem", "Dima", 10)

	blockchain.AddBlock()

	blockchain.Log()

}
