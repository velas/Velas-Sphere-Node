package handler

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/syndtr/goleveldb/leveldb"
	ethdepositcontract "github.com/velas/Velas-Sphere-Contracts/ethdepositcontract"
)

type Config struct {
	DB                         *leveldb.DB
	Contract                   *ethdepositcontract.Ethdepositcontract
	TransactOptionsInitializer func(key *ecdsa.PrivateKey) (*bind.TransactOpts, error)
}
