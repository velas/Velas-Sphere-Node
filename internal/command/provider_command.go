package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewProviderCommand() *cobra.Command {
	return &cobra.Command{
		Use: "provider",
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
