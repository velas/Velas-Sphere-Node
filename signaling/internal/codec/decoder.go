package codec

import (
	"encoding/binary"
	"fmt"
	"signaling/internal/entity"
	"signaling/internal/enum"
)

func DecodeMessage(messageBytes []byte) (entity.Message, error) {
	if len(messageBytes) < 32 {
		return entity.Message{}, fmt.Errorf("invalid message length")
	}
	message := entity.Message{}
	message.Version = binary.LittleEndian.Uint64(messageBytes[:8])
	switch message.Version {
	case 1:
		message.ID = binary.LittleEndian.Uint64(messageBytes[8:16])
		message.Answer = binary.LittleEndian.Uint64(messageBytes[16:24])
		message.Kind = enum.MessageKind(binary.LittleEndian.Uint64(messageBytes[24:32]))

		switch message.Kind {
		case enum.UploadInitializationMessageKind:
			var err error
			message.Body, err = DecodeInitUploadMessage(messageBytes[32:])
			if err != nil {
				return entity.Message{}, fmt.Errorf("failed to decode the init upload message")
			}

		case enum.BlockTransmissionMessageKind:
			var err error
			message.Body, err = DecodeBlockTransmissionMessage(messageBytes[32:])
			if err != nil {
				return entity.Message{}, fmt.Errorf("failed to decode the init upload message")
			}

		case enum.DownloadInitializationMessageKind:
			var err error
			message.Body, err = DecodeInitDownloadMessage(messageBytes[32:])
			if err != nil {
				return entity.Message{}, fmt.Errorf("failed to decode the init upload message")
			}
		}

		return message, nil

	default:
		return entity.Message{}, fmt.Errorf("unsupported message version")
	}
}

func DecodeInitUploadMessage(messageBytes []byte) (entity.UploadInitializationMessage, error) {
	if len(messageBytes) < 16 {
		return entity.UploadInitializationMessage{}, fmt.Errorf("invalid upload message length")
	}
	hash := messageBytes[:8]
	size := binary.LittleEndian.Uint64(messageBytes[8:16])
	return entity.UploadInitializationMessage{
		Hash: hash,
		Size: size,
	}, nil
}

func DecodeBlockTransmissionMessage(messageBytes []byte) (entity.BlockTransmissionMessage, error) {
	if len(messageBytes) < 8 {
		return entity.BlockTransmissionMessage{}, fmt.Errorf("invalid upload message length")
	}
	offset := binary.LittleEndian.Uint64(messageBytes[0:8])
	return entity.BlockTransmissionMessage{
		Offset: offset,
		Data:   messageBytes[8:],
	}, nil
}

func DecodeInitDownloadMessage(messageBytes []byte) (entity.DownloadInitializationMessage, error) {
	if len(messageBytes) < 24 {
		return entity.DownloadInitializationMessage{}, fmt.Errorf("invalid upload message length")
	}
	hash := messageBytes[:8]
	fromOffset := binary.LittleEndian.Uint64(messageBytes[8:16])
	toOffset := binary.LittleEndian.Uint64(messageBytes[16:24])
	return entity.DownloadInitializationMessage{
		Hash:       hash,
		FromOffset: fromOffset,
		ToOffset:   toOffset,
	}, nil
}
