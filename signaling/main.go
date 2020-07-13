package main

import (
	"encoding/json"
	"fmt"
	"signaling/internal/entity"
	"signaling/internal/server"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

func bsonPoC() {
	data := []byte{}
	for i := byte(0); i < 255; i++ {
		data = append(data, i)
	}

	myMessage := entity.Message{
		ID:   1,
		Kind: 2,
		Body: entity.BlockTransmissionMessage{
			Offset: 16,
			Data:   data,
		},
		Answer: 0,
	}

	jsonBytes, _ := json.Marshal(myMessage)

	fmt.Println(len(jsonBytes))

	bsonBytes, _ := bson.Marshal(myMessage)

	fmt.Println(len(bsonBytes))

}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(*sync.WaitGroup) {
		defer wg.Done()
		server.FileServer()
	}(&wg)

	wg.Add(1)
	go func(*sync.WaitGroup) {
		defer wg.Done()
		server.SignalingServer()
	}(&wg)

	wg.Wait()
}
