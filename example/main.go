package main

import (
	"github.com/xamber/capusta"
	"log"
)

var blockchain = capusta.Blockchain

func main() {

	blockchain.MineBlock("Artem")

	_, err := blockchain.TransferMoney("Artem", "Dima", 100)
	if err != nil {
		log.Fatal(err)
	}

	blockchain.MineBlock("Dima")

	_, err = blockchain.TransferMoney("Dima", "Artem", 50)
	if err != nil {
		log.Fatal(err)
	}

	blockchain.MineBlock("Artem")

	_, err = blockchain.TransferMoney("Artem", "Dima", 80)
	if err != nil {
		log.Fatal(err)
	}

	blockchain.MineBlock("Dima")
	_, err = blockchain.TransferMoney("Dima", "Artem", 77)
	if err != nil {
		log.Fatal(err)
	}
	_, err = blockchain.TransferMoney("Artem", "Dima", 10)
	if err != nil {
		log.Fatal(err)
	}

	blockchain.MineBlock("Artem")

	blockchain.Info()

}
