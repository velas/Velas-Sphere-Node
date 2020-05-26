package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

func NewGetFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		file, err := os.Open(id)
		if err != nil {
			log.Println("failed to read the file:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.ServeContent(w, r, id, time.Now(), file)
	}
}
