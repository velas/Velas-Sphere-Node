package server

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gabriel-vasile/mimetype"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/entropy"
	"github.com/sorenvonsarvort/velas-sphere/internal/merkletree"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/syndtr/goleveldb/leveldb"
)

type StorageServer struct {
	DB                 *leveldb.DB
	Ethdepositcontract *contract.Ethdepositcontract
}

func (s StorageServer) SaveFile(ctx context.Context, req *resources.FileStorageRequest) (*resources.FileStorageResponse, error) {
	if s.DB == nil {
		return nil, errors.New("no storage provided")
	}

	db := s.DB

	name := req.GetName()
	if name == "" {
		return nil, errors.New("the name cannot be empty")
	}

	// TODO: add contract data to the request and check pricing

	data := req.GetData()

	mime, err := mimetype.DetectReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to detect mime type of the file: %w", err)
	}
	if !mime.Is("application/octet-stream") {
		return nil, errors.New("the file has invalid mime type")
	}

	// 7.99 bits per byte is a heuristic threshold
	// 2.99 is used for testing
	e := entropy.Shannon(data)
	if e <= 2.99 {
		return nil, errors.New("the file entropy is too low, consider larger file and ensure encryption")
	}

	merkleTreeRootBytes := req.GetMerkleTreeRoot()

	merkleTreeRoot := merkletree.Node{}
	err = json.Unmarshal(merkleTreeRootBytes, &merkleTreeRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the merkle tree: %w", err)
	}

	treeIsValid := merkletree.VerifyNode(merkleTreeRoot)
	if !treeIsValid {
		return nil, fmt.Errorf("invalid tree provided")
	}

	requesterPublicKey := req.GetRequesterPublicKey()

	id := hex.EncodeToString(requesterPublicKey) + "/" + req.GetName()

	getBackToken := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, getBackToken)
	if err != nil {
		return nil, fmt.Errorf("failed to generate a get back token: %w", err)
	}

	verificationToken := make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, verificationToken)
	if err != nil {
		return nil, fmt.Errorf("failed to generate a verification token: %w", err)
	}

	storedFileJSONBytes, err := json.Marshal(
		entity.File{
			RequesterPublicKey: requesterPublicKey,
			Data:               data,
			MerkleTreeRoot:     merkleTreeRoot,
			GetBackToken:       getBackToken,
			VerificationToken:  verificationToken,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the stored file: %w", err)
	}

	err = db.Put([]byte("stored_files/"+id), storedFileJSONBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to store the file: %w", err)
	}

	return &resources.FileStorageResponse{
		Id:                id,
		GetBackToken:      getBackToken,
		VerificationToken: verificationToken,
	}, nil
}

func (s StorageServer) GetFileBack(ctx context.Context, req *resources.GetFileBackRequest) (*resources.GetFileBackResponse, error) {
	if s.DB == nil {
		return nil, errors.New("no storage provided")
	}
	db := s.DB

	if s.Ethdepositcontract == nil {
		return nil, errors.New("no contract instance injected")
	}
	ethDepositContract := s.Ethdepositcontract

	id := req.GetId()

	rawFile, err := db.Get([]byte("stored_files/"+id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get the file: %w", err)
	}

	file := entity.File{}
	err = json.Unmarshal(rawFile, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the file: %w", err)
	}

	// TODO: why do we compare token with it's signature??? We need to verify it using the public key
	if bytes.Compare(req.GetGetBackTokenSignature(), file.GetBackToken) != 0 {
		return nil, errors.New("invalid get back token signature")
	}

	invoiceTx, err := ethDepositContract.CreateInvoice(nil, nil, nil, common.Address{}, nil, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("invoice creation failed: %w", err)
	}

	return &resources.GetFileBackResponse{
		Id:        id,
		InvoiceID: invoiceTx.Hash().String(), // TODO: replace by a real invoice id
		Data:      file.Data,
	}, nil
}

func (s StorageServer) VerifyFileStorage(ctx context.Context, req *resources.FileStorageVerificationRequest) (*resources.FileStorageVerificationResponse, error) {
	if s.DB == nil {
		return nil, errors.New("no storage provided")
	}

	db := s.DB

	id := req.GetId()

	rawFile, err := db.Get([]byte("stored_files/"+id), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get the file: %w", err)
	}

	file := entity.File{}
	err = json.Unmarshal(rawFile, &file)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal the file: %w", err)
	}

	verificationTokenHash := crypto.Keccak256Hash(file.VerificationToken)

	sigPublicKey, err := crypto.Ecrecover(verificationTokenHash.Bytes(), req.GetVerificationTokenSignature())
	if err != nil {
		return nil, fmt.Errorf("filed to recover the public key: %w", err)
	}

	matches := bytes.Equal(sigPublicKey, file.RequesterPublicKey)

	// TODO: why do we compare token with it's signature??? We need to verify it using the public key
	if !matches {
		return nil, errors.New("invalid verification token signature")
	}

	// TODO: store last verification time?
	// TODO: if got no verification requests during some time, generate invoice and set a deadline for file removal?

	x := sha256.New()
	x.Write(append(file.Data, req.GetChallenge()...))

	path, err := merkletree.FindPath(file.MerkleTreeRoot, x.Sum(nil))
	if err != nil {
		return nil, fmt.Errorf("failed to find the path: %w", err)
	}

	file.VerificationToken = []byte("new_verification_token")

	updatedFileJSONBytes, err := json.Marshal(file)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the updated file: %w", err)
	}

	err = db.Put([]byte("stored_files/"+id), updatedFileJSONBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to store the updated file: %w", err)
	}

	return &resources.FileStorageVerificationResponse{
		Id:                   req.GetId(),
		Path:                 strings.Join(path, ":"),
		NewVerificationToken: file.VerificationToken,
	}, nil
}
