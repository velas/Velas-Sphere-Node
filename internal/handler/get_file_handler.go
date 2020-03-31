package handler

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

func NewGetFileHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.DB == nil {
			log.Println("no db provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		db := config.DB

		request := entity.File{}
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read body:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(requestBody, &request)
		if err != nil {
			log.Println("failed to unmarshal json:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		rawEntry, err := db.Get([]byte("requested_files/"+request.ID), nil)
		if err != nil {
			log.Println("failed to get the entry")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		entry := entity.File{}

		err = json.Unmarshal(rawEntry, &entry)
		if err != nil {
			log.Println("failed to unmarshal entry:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		conn, err := grpc.Dial(request.Target, grpc.WithInsecure())
		if err != nil {
			log.Println("failed to dial the target")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		c := resources.NewStorageClient(conn)

		response, err := c.GetFileBack(
			context.TODO(),
			&resources.GetFileBackRequest{
				Id:                    request.ID,
				GetBackTokenSignature: entry.GetBackToken,
			},
		)
		if err != nil {
			log.Println("failed to request the file back:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response.Data, err = decrypt([]byte("myawesomekey0000"), response.Data)
		if err != nil {
			log.Println("failed to decrypt the file back:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseJSONBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("failed to render the response:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(responseJSONBytes)
	}
}
