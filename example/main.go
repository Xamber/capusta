package main

import (
	"github.com/xamber/capusta"
)

var blockchain = capusta.Blockchain

func main() {

	blockchain.MineBlock("Artem")

	blockchain.AddTransaction("Artem", "Dima", 100)
	blockchain.AddTransaction("Dima", "Artem", 50)
	blockchain.AddTransaction("Artem", "Dima", 80)

	blockchain.MineBlock("Dima")

	blockchain.AddTransaction("Dima", "Artem", 77)
	blockchain.AddTransaction("Artem", "Dima", 10)

	blockchain.MineBlock("Artem")

	blockchain.Info()

}
