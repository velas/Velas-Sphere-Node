package handler

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	// TODO: inject verification tasks channel
	DB *leveldb.DB
}
