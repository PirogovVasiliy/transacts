package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type counter struct {
	Value int `json:"value"`
}

func ListenEvents(ctx context.Context, network *client.Network) {
	fmt.Println("--- Начинаем слушать события! ---")
	fmt.Println("--------------------------------------")

	events, err := network.ChaincodeEvents(ctx, chaincodeName)
	if err != nil {
		log.Fatalln("Ошибка прослушивания событий!", err)
	}

	go func() {
		for event := range events {
			var counter counter
			err := json.Unmarshal(event.Payload, &counter)
			if err != nil {
				log.Println("Ошибка получения события!")
			}
			fmt.Println("--- Получено событие! ---")
			fmt.Println("Название:", event.EventName)
			fmt.Println("Значение:", counter.Value)
			fmt.Println("--------------------------------------")
		}
	}()

}
