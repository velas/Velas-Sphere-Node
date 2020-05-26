package handler

import (
	"crypto/ecdsa"

	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	DB                  *leveldb.DB
	ContractInitializer func(key *ecdsa.PrivateKey) (*contract.Ethdepositcontract, error)
}
