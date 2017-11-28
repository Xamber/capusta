package main

import (
	"fmt"
	"github.com/xamber/capusta"
)

var blockchain = capusta.Blockchain

func main() {

	artem := blockchain.GetWallet("Artem")
	dima := blockchain.GetWallet("Dima")

	blockchain.MineBlock("Artem")
	artem.TransferMoney(dima,100)

	blockchain.MineBlock("Dima")
	dima.TransferMoney(artem, 50)

	blockchain.MineBlock("Artem")
	artem.TransferMoney(dima,80)

	blockchain.MineBlock("Dima")

	dima.TransferMoney(artem, 77)
	artem.TransferMoney(dima,10)

	blockchain.MineBlock("Artem")

	fmt.Println(blockchain)
}
