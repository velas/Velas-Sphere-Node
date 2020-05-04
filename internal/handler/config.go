package handler

import (
	"crypto/ecdsa"

	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	DB         *leveldb.DB
	PrivateKey *ecdsa.PrivateKey
}
