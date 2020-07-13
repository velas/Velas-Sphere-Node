package codec

import (
	"encoding/binary"
	"fmt"
	"signaling/internal/entity"
	"signaling/internal/enum"
)

func EncodeMessage(message entity.Message) ([]byte, error) {
	switch message.Version {
	case 1:
		result := []byte{}

		buf := make([]byte, 8)
		binary.LittleEndian.PutUint64(buf, message.Version)
		result = append(result, buf...)

		binary.LittleEndian.PutUint64(buf, message.ID)
		result = append(result, buf...)

		binary.LittleEndian.PutUint64(buf, message.Answer)
		result = append(result, buf...)

		switch body := message.Body.(type) {
		case entity.UploadInitializationMessage:
			binary.LittleEndian.PutUint64(buf, uint64(enum.UploadInitializationMessageKind))
			result = append(result, buf...)

			encodedBody, err := EncodeInitUploadMessage(body)
			if err != nil {
				return nil, fmt.Errorf("failed to encode the init upload message: %w", err)
			}

			result = append(result, encodedBody...)

		case entity.BlockTransmissionMessage:
			binary.LittleEndian.PutUint64(buf, uint64(enum.BlockTransmissionMessageKind))
			result = append(result, buf...)

			encodedBody, err := EncodeBlockTransmissionMessage(body)
			if err != nil {
				return nil, fmt.Errorf("failed to encode the init upload message: %w", err)
			}

			result = append(result, encodedBody...)

		case entity.DownloadInitializationMessage:
			binary.LittleEndian.PutUint64(buf, uint64(enum.DownloadInitializationMessageKind))
			result = append(result, buf...)

			encodedBody, err := EncodeInitDownloadMessage(body)
			if err != nil {
				return nil, fmt.Errorf("failed to encode the init upload message: %w", err)
			}

			result = append(result, encodedBody...)

		default:
			return nil, fmt.Errorf("unsupported message body type")
		}

		return result, nil
	}
	return nil, fmt.Errorf("unsupported message version")
}

func EncodeInitUploadMessage(message entity.UploadInitializationMessage) ([]byte, error) {
	result := message.Hash

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, message.Size)
	result = append(result, buf...)

	return result, nil
}

func EncodeInitDownloadMessage(message entity.DownloadInitializationMessage) ([]byte, error) {
	result := message.Hash

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, message.FromOffset)
	result = append(result, buf...)

	binary.LittleEndian.PutUint64(buf, message.ToOffset)
	result = append(result, buf...)

	return result, nil
}

func EncodeBlockTransmissionMessage(message entity.BlockTransmissionMessage) ([]byte, error) {
	result := make([]byte, 8)

	binary.LittleEndian.PutUint64(result, message.Offset)

	result = append(result, message.Data...)

	return result, nil
}
