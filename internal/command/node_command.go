package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/initializer"
	"github.com/sorenvonsarvort/velas-sphere/internal/service"
	"github.com/spf13/cobra"
	"github.com/syndtr/goleveldb/leveldb"
)

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
			if err == nil {
				log.Println("found a config")

				err = json.Unmarshal(configBytes, &config)
				if err != nil {
					return fmt.Errorf("failed to parse the config: %w", err)
				}
			}

			client, err := ethclient.Dial(config.Node.EthereumNodeTarget)
			if err != nil {
				log.Fatal(err)
			}

			contractBuilder := initializer.ContractInitializer(client, config.Node.EthdepositcontractAddress)

			db, err := leveldb.OpenFile("db", nil)
			if err != nil {
				return errors.Wrap(err, "failed to open the db file")
			}
			defer db.Close()

			wg := sync.WaitGroup{}

			mux := service.Mux(
				&wg,
				map[string]service.Service{
					"storage":   service.Storage(db, contractBuilder),
					"streaming": service.Streaming,
				},
			)
			err = mux()
			if err != nil {
				log.Println("mux service failed:", err)
			}

			wg.Wait()

			return errors.New("all services are stopped")
		},
	}
}
