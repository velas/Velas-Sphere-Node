package handler

import (
	"log"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"

	"github.com/sorenvonsarvort/velas-sphere/internal/enum"
)

func NewGetNodesStatisticHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.DB == nil {
			log.Println("no db provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		db := config.DB

		entry, err := db.Get([]byte(enum.NodesStatisticKey), nil)
		if err != nil {
			if err == leveldb.ErrNotFound {
				w.Write([]byte(`{"nodes_count":0, "nodes_space":0}`))
				return
			}
			log.Println("failed to get the node statistic entry")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(entry)
	}
}
