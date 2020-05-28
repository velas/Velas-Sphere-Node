package service

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sorenvonsarvort/velas-sphere/internal/handler"
)

func Storage(handlerConfig handler.Config) Service {
	return func() error {
		r := chi.NewRouter()

		r.Use(
			middleware.SetHeader(
				"Content-Type",
				"application/json",
			),
		)

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
