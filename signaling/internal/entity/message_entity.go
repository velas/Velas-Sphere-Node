package entity

import "signaling/internal/enum"

// From client perspective
// TODO: datalimit

type Message struct {
	Version uint64
	ID      uint64
	Answer  uint64
	Kind    enum.MessageKind
	Body    interface{}
}

type UploadInitializationMessage struct {
	Hash []byte
	Size uint64
}

type DownloadInitializationMessage struct {
	Hash       []byte
	FromOffset uint64 // 0 = from start
	ToOffset   uint64 // 0 = up to end
}

type BlockTransmissionMessage struct {
	Offset uint64
	Data   []byte
}

// TODO: heartbeat message
