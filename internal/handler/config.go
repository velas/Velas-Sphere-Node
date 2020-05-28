package handler

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	DB                         *leveldb.DB
	Contract                   *contract.Ethdepositcontract
	TransactOptionsInitializer func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error)
}
