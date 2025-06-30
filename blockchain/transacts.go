package blockchain

import (
	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func CallCreateCounter(contract *client.Contract, value string) error {
	_, err := contract.SubmitTransaction("CreateCounter", value)
	if err != nil {
		return err
	}
	return nil
}

func CallReadCounter(contract *client.Contract) error {
	_, err := contract.EvaluateTransaction("ReadCounter")
	if err != nil {
		return err
	}
	return nil
}

func CallIncrimentCounter(contract *client.Contract, value string) error {
	_, err := contract.SubmitTransaction("IncrimentCounter", value)
	if err != nil {
		return err
	}
	return nil
}

func CallMinusCounter(contract *client.Contract) error {
	_, err := contract.SubmitTransaction("MinusCounter")
	if err != nil {
		return err
	}
	return nil
}
