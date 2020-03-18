package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/enum"
	"github.com/sorenvonsarvort/velas-sphere/internal/handler"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

func NewNodeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "node",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := Config{
				Provider: ProviderConfig{
					PluginTarget: "plugin:8082",
				},
			}

			configBytes, err := ioutil.ReadFile("config.json")
			if err == nil {
				log.Println("found config")
				json.Unmarshal(configBytes, &config)
			}

			db, err := leveldb.OpenFile("db", nil)
			if err != nil {
				return errors.Wrap(err, "failed to open the db file")
			}
			defer db.Close()

			initialEntry, err := json.Marshal(
				entity.NodeInfo{
					Name:    "Dev Node",
					Version: "0.0.1",
				},
			)
			if err != nil {
				return errors.Wrap(err, "failed to marshal the initial entry")
			}

			err = db.Put([]byte(enum.NodeInfoKey), initialEntry, nil)
			if err != nil {
				return errors.Wrap(err, "failed to perform initial write")
			}

			r := chi.NewRouter()

			r.Use(
				middleware.SetHeader(
					"Content-Type",
					"application/json",
				),
			)

			r.Get(
				"/info",
				handler.NewGetInfoHandler(
					handler.Config{
						DB: db,
					},
				),
			)

			r.Post(
				"/task_execution_request",
				handler.NewPostTaskExecutionRequestHandler(
					handler.Config{
						DB: db,
					},
				),
			)

			log.Println("api service started")

			go http.ListenAndServe(":3000", r)

			lis, err := net.Listen("tcp", ":8081")
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			conn, err := grpc.Dial(config.Provider.PluginTarget, grpc.WithInsecure())
			if err != nil {
				return errors.Wrap(err, "failed to dial the plugin")
			}
			defer conn.Close()

			s := grpc.NewServer()
			resources.RegisterProviderServer(
				s,
				server.ProviderServer{
					PluginClient: resources.NewPluginClient(conn),
				},
			)

			log.Println("provider service started")

			return s.Serve(lis)
		},
	}
}
