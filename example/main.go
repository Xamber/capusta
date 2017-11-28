package main

import (
	"fmt"
	"github.com/xamber/capusta"
	"log"
)

var blockchain = capusta.Blockchain

func main() {

	printTransaction := func(id string, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(id)
		}
	}

	artem := blockchain.GetWallet("Artem")
	dima := blockchain.GetWallet("Dima")

	blockchain.MineBlock("Artem")
	fmt.Println(artem.GetBalance())

	id, err := artem.TransferMoney(dima,100)
	printTransaction(id, err)

	blockchain.MineBlock("Dima")

	id, err = dima.TransferMoney(artem, 50)
	printTransaction(id, err)

	blockchain.MineBlock("Artem")

	id, err = artem.TransferMoney(dima,80)
	printTransaction(id, err)

	blockchain.MineBlock("Dima")

	id, err = dima.TransferMoney(artem, 77)
	printTransaction(id, err)
	id, err = artem.TransferMoney(dima,10)
	printTransaction(id, err)

	blockchain.MineBlock("Artem")

	fmt.Println(blockchain)
}
