package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
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

// TODO: run file service

func NewNodeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "node",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := Config{
				Node: NodeConfig{
					PluginTarget:        "plugin:8082",
					StoragePluginTarget: "storage:8083",
					EthereumNodeTarget:  "http://127.0.0.1:8545",
					// TODO: price tolarance config
					// TODO: keypair config
				},
			}

			client, err := ethclient.Dial(config.Node.EthereumNodeTarget)
			if err != nil {
				log.Fatal(err)
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
				entity.Info{
					Name:    "Dev Node",
					Version: "0.0.1",
					// TODO: price tolerance info injection
					// TODO: supported plugins info injection
					// TODO: public key info injection
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
				"/task",
				handler.NewPostTaskHandler(
					handler.Config{
						DB: db,
					},
				),
			)

			r.Post(
				"/file",
				handler.NewPostFileHandler(
					handler.Config{
						DB: db,
					},
				),
			)

			wg := sync.WaitGroup{}
			wg.Add(3)

			go func(wg *sync.WaitGroup) {
				log.Println("api service started")

				defer func() {
					// We do not need to handle any panics, since it's already done by the router!
					wg.Done()
				}()

				err := http.ListenAndServe(":3000", r)
				if err != nil {
					log.Println("failed to listen and serve:", err)
				}
			}(&wg)

			lis, err := net.Listen("tcp", ":8081")
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			conn, err := grpc.Dial(config.Node.PluginTarget, grpc.WithInsecure())
			if err != nil {
				return errors.Wrap(err, "failed to dial the plugin")
			}
			defer conn.Close()

			s := grpc.NewServer()
			resources.RegisterProviderServer(
				s,
				server.ProviderServer{
					PluginClient: resources.NewPluginClient(conn),
					EthClient:    client,
				},
			)

			go func(wg *sync.WaitGroup) {
				log.Println("provider service started")

				defer func() {
					r := recover()
					if r != nil {
						log.Println("got panic:", r)
					}

					wg.Done()
				}()

				err := s.Serve(lis)
				if err != nil {
					log.Println("failed to serve:", err)
				}
			}(&wg)

			storageListener, err := net.Listen("tcp", ":8083")
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			storageServer := grpc.NewServer()
			resources.RegisterStorageServer(storageServer, server.StorageServer{})

			go func(wg *sync.WaitGroup) {
				log.Println("storage service started")

				defer func() {
					r := recover()
					if r != nil {
						log.Println("got panic:", r)
					}

					wg.Done()
				}()

				err := storageServer.Serve(storageListener)
				if err != nil {
					log.Println("failed to serve:", err)
				}
			}(&wg)

			wg.Wait()

			return errors.New("all services are stopped")
		},
	}
}
