package server

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/entropy"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type StorageServer struct {
	// TODO: inject leveldb client
	Storage map[string]entity.File
}

func (s StorageServer) SaveFile(ctx context.Context, req *resources.FileStorageRequest) (*resources.FileStorageResponse, error) {
	// TODO: add contract data to the request and check pricing

	data := req.GetData()

	reader := strings.NewReader(data) // ignoring error for brevity's sake
	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to detect mime type of the file: %w", err)
	}
	if mime.String() != "AES" {
		return nil, errors.New("the file has invalid mime type")
	}

	e := entropy.Shannon(data)
	// 0.799 is a heuristic AES threshold
	if e <= 0.799 {
		return nil, errors.New("the file entropy is too low")
	}

	// TODO: validate merkle tree
	// TODO: use public key as suffix or other meta for id
	// TODO: decode data
	// TODO: validate public key
	// TODO: validate signature

	id := "some_prefix" + req.GetName()

	// TODO: store file in the leveldb
	// TODO: include proper data
	s.Storage[id] = entity.File{
		RequesterPublicKey: req.GetRequesterPublicKey(),
		Data:               []byte(req.GetData()),
		MerkleTreeRoot:     req.GetMerkleTreeRoot(),
		GetBackToken:       "rand",
	}

	// TODO: generate the get back token
	// TODO: generate the verification token
	getBackToken := "get_back_token"
	verificationToken := "verification_token"

	return &resources.FileStorageResponse{
		Id:                id,
		GetBackToken:      getBackToken,
		VerificationToken: verificationToken,
	}, nil
}

func (s StorageServer) GetFileBack(ctx context.Context, req *resources.GetFileBackRequest) (*resources.GetFileBackResponse, error) {
	return nil, nil
}

func (s StorageServer) VerifyFileStorage(ctx context.Context, req *resources.FileStorageVerificationRequest) (*resources.FileStorageVerificationResponse, error) {
	// TODO: if got no verification requests during some time, generate invoice and set a deadline for file removal?
	return nil, nil
}
