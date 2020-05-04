package command

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/contract"
	"github.com/sorenvonsarvort/velas-sphere/internal/entity"
	"github.com/sorenvonsarvort/velas-sphere/internal/enum"
	"github.com/sorenvonsarvort/velas-sphere/internal/handler"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/grpc"
)

func initContract(ethClient *ethclient.Client, key *ecdsa.PrivateKey, address common.Address) (*contract.Ethdepositcontract, error) {
	publicKey := key.PublicKey

	fromAddress := crypto.PubkeyToAddress(publicKey)
	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested gas price: %w", err)
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	c, err := contract.NewEthdepositcontract(address, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create new deposit contract: %w", err)
	}

	return c, nil
}

func NewNodeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "node",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := Config{
				Node: NodeConfig{
					PluginTarget:       "plugin:8082",
					EthereumNodeTarget: "http://127.0.0.1:8545",
					// TODO: price tolarance config
				},
			}

			configBytes, err := ioutil.ReadFile("config.json")
			if err != nil {
				return fmt.Errorf("failed to read the config: %w", err)
			}

			err = json.Unmarshal(configBytes, &config)
			if err != nil {
				return fmt.Errorf("failed to parse the config: %w", err)
			}

			keystoreBytes, err := ioutil.ReadFile(config.Node.KeystoreFilePath)
			if err != nil {
				return fmt.Errorf("failed to read the keystore file: %w", err)
			}

			k, err := keystore.DecryptKey(
				keystoreBytes,
				string(config.Node.KeystorePassword),
			)
			if err != nil {
				return fmt.Errorf("failed to decrypt the keystore: %w", err)
			}
			privateKey := k.PrivateKey

			client, err := ethclient.Dial(config.Node.EthereumNodeTarget)
			if err != nil {
				log.Fatal(err)
			}

			ethDepositContract, err := initContract(client, privateKey, config.Node.EthdepositcontractAddress)
			if err != nil {
				log.Fatal(err)
			}

			// publicKey := privateKey.Public()
			// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			// if !ok {
			// log.Fatal("error casting public key to ECDSA")
			// }

			// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			// TODO: check node registration

			// registerTx, err := ethDepositContract.RegisterNode(nil, fromAddress, nil, nil, nil, nil)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// _ = registerTx

			// fmt.Println("registered node:", registerTx.Hash())

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

			handlerConfig := handler.Config{
				DB:         db,
				PrivateKey: privateKey,
			}

			r.Get(
				"/info",
				handler.NewGetInfoHandler(
					handlerConfig,
				),
			)

			r.Get(
				"/file",
				handler.NewGetFileHandler(
					handlerConfig,
				),
			)

			r.Post(
				"/task",
				handler.NewPostTaskHandler(
					handlerConfig,
				),
			)

			r.Post(
				"/file",
				handler.NewPostFileHandler(
					handlerConfig,
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
					PluginClient:       resources.NewPluginClient(conn),
					Ethdepositcontract: ethDepositContract,
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
			resources.RegisterStorageServer(
				storageServer,
				server.StorageServer{
					// TODO: leveldb injection
					DB: db},
			)

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
