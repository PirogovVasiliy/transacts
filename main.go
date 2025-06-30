package main

import (
	"fmt"
	"log"
	"transactsender/blockchain"
)

func main() {

	contract := blockchain.ConnectToContract()

	fmt.Println("--- Submit CreateCounter!")
	err := blockchain.CallCreateCounter(contract, "5")
	if err != nil {
		log.Fatalln("Ошибка CreateCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	fmt.Println("--- Submit ReadCounter!")
	err = blockchain.CallReadCounter(contract)
	if err != nil {
		log.Fatalln("Ошибка ReadCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	fmt.Println("--- Submit IncrimentCounter!")
	err = blockchain.CallIncrimentCounter(contract, "4")
	if err != nil {
		log.Fatalln("Ошибка IncrimentCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	fmt.Println("--- Submit ReadCounter!")
	err = blockchain.CallReadCounter(contract)
	if err != nil {
		log.Fatalln("Ошибка ReadCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	fmt.Println("--- Submit MinusCounter!")
	err = blockchain.CallMinusCounter(contract)
	if err != nil {
		log.Fatalln("Ошибка MinusCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")

	fmt.Println("--- Submit ReadCounter!")
	err = blockchain.CallReadCounter(contract)
	if err != nil {
		log.Fatalln("Ошибка ReadCounter", err)
	}
	fmt.Println("*** Transaction committed successfully!")
	fmt.Println("--------------------------------------")
}
