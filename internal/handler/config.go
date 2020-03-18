package handler

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type Config struct {
	DB *leveldb.DB
}
