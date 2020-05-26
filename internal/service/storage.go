package service

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/sorenvonsarvort/velas-sphere/internal/handler"
	"github.com/syndtr/goleveldb/leveldb"
)

func Storage(db *leveldb.DB, contractInitializer func(key *ecdsa.PrivateKey) (*contract.Ethdepositcontract, error)) Service {
	return func() error {
		r := chi.NewRouter()

		r.Use(
			middleware.SetHeader(
				"Content-Type",
				"application/json",
			),
		)

		handlerConfig := handler.Config{
			DB:                  db,
			ContractInitializer: contractInitializer,
		}

		r.Get(
			"/file/{id}",
			handler.NewGetFileHandler(
				handlerConfig,
			),
		)

		r.Post(
			"/file",
			handler.NewPostFileHandler(
				handlerConfig,
			),
		)

		err := http.ListenAndServe(":3000", r)
		if err != nil {
			return fmt.Errorf("failed to listen and serve: %w", err)
		}

		return nil
	}
}
