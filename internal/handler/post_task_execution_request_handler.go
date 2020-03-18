package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/enum"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

func NewPostTaskExecutionRequestHandler(config Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.DB == nil {
			log.Println("no db provided")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		db := config.DB

		requestBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read the request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		request := entity.TaskExecutionRequest{}

		err = json.Unmarshal(requestBytes, &request)
		if err != nil {
			log.Println("failed to unmarshal the request body")
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

		c := resources.NewProviderClient(conn)

		providerResponse, err := c.RequestTaskExecution(r.Context(), &resources.TaskExecutionRequest{
			Id:    request.ID,
			Input: request.Input,
		})
		if err != nil {
			log.Println("failed to request task execution")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = db.Put([]byte(enum.LastTaskKey), []byte(providerResponse.String()), nil)
		if err != nil {
			log.Println("failed to save the task")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		providerResponseJSONBytes, err := json.Marshal(providerResponse)
		if err != nil {
			log.Println("failed to marshal the provider response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(providerResponseJSONBytes)
	}
}
