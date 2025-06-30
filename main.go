package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"transactsender/blockchain"
)

func main() {

	network, contract := blockchain.ConnectToContract()
	fmt.Println("--- Успешное подключение к контракту! ---")
	fmt.Println("--------------------------------------")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	blockchain.ListenEvents(ctx, network)

	time.Sleep(time.Second)

	fmt.Println("--- Submit CreateCounter!")
	err := blockchain.CallCreateCounter(contract, "5")
	if err != nil {
		log.Fatalln("Ошибка CreateCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	time.Sleep(2 * time.Second)

	fmt.Println("--- Submit IncrimentCounter!")
	err = blockchain.CallIncrimentCounter(contract, "4")
	if err != nil {
		log.Fatalln("Ошибка IncrimentCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	time.Sleep(3 * time.Second)

	fmt.Println("--- Submit MinusCounter!")
	err = blockchain.CallMinusCounter(contract)
	if err != nil {
		log.Fatalln("Ошибка MinusCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	select {}
}
