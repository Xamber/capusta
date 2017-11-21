package main

import (
	"github.com/xamber/capusta"
)

var blockchain = capusta.Blockchain

func main() {

	blockchain.MineBlock("Artem")

	blockchain.TransferMoney("Artem", "Dima", 100)
	blockchain.TransferMoney("Dima", "Artem", 50)
	blockchain.TransferMoney("Artem", "Dima", 80)

	blockchain.MineBlock("Dima")

	blockchain.TransferMoney("Dima", "Artem", 77)
	blockchain.TransferMoney("Artem", "Dima", 10)

	blockchain.MineBlock("Artem")

	blockchain.Info()

}
