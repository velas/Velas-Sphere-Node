package main

import (
	"fmt"
	"signaling/internal/codec"
	"signaling/internal/entity"
)

func initUploadMessageDemo() {
	fmt.Println("init upload message demo")
	msg := entity.Message{
		Version: 1,
		ID:      2,
		Answer:  3,
		Body: entity.UploadInitializationMessage{
			Hash: []byte{1, 2, 3, 4, 5, 6, 7, 8},
			Size: 4,
		},
	}
	encoded, err := codec.EncodeMessage(msg)
	if err != nil {
		fmt.Println("failed to encode the message:", err)
	}

	decoded, err := codec.DecodeMessage(encoded)
	if err != nil {
		fmt.Println("failed to decode the message:", err)
	}

	fmt.Printf("%#v\n", decoded)
}

func initInitDownloadMessage() {
	fmt.Println("init download message demo")
	msg := entity.Message{
		Version: 1,
		ID:      2,
		Answer:  3,
		Body: entity.DownloadInitializationMessage{
			Hash:       []byte{1, 2, 3, 4, 5, 6, 7, 8},
			FromOffset: 4,
			ToOffset:   48,
		},
	}
	encoded, err := codec.EncodeMessage(msg)
	if err != nil {
		fmt.Println("failed to encode the message:", err)
	}

	decoded, err := codec.DecodeMessage(encoded)
	if err != nil {
		fmt.Println("failed to decode the message:", err)
	}

	fmt.Printf("%#v\n", decoded)
}

func blockTransmissionMessage() {
	fmt.Println("download block message demo")
	msg := entity.Message{
		Version: 1,
		ID:      2,
		Answer:  3,
		Body: entity.BlockTransmissionMessage{
			Offset: 48,
			Data:   []byte{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	encoded, err := codec.EncodeMessage(msg)
	if err != nil {
		fmt.Println("failed to encode the message:", err)
	}

	decoded, err := codec.DecodeMessage(encoded)
	if err != nil {
		fmt.Println("failed to decode the message:", err)
	}

	fmt.Printf("%#v\n", decoded)
}

func main() {
	initUploadMessageDemo()
	initInitDownloadMessage()
	blockTransmissionMessage()
}
