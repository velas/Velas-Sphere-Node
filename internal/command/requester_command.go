package command

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewRequesterCommand() *cobra.Command {
	return &cobra.Command{
		Use: "requester",
		RunE: func(cmd *cobra.Command, args []string) error {
			config := Config{
				RequesterConfig: RequesterConfig{
					Target: "provider:8081",
				},
			}

			configBytes, err := ioutil.ReadFile("config.json")
			if err == nil {
				log.Println("found config")
				json.Unmarshal(configBytes, &config)
			}

			log.Println("requester started")

			conn, err := grpc.Dial(config.RequesterConfig.Target, grpc.WithInsecure())
			if err != nil {
				return errors.Wrap(err, "failed to dial the provider")
			}
			defer conn.Close()

			c := resources.NewProviderClient(conn)

			for {
				select {
				case <-cmd.Context().Done():
					return cmd.Context().Err()
				default:
				}

				resp, err := c.RequestTaskExecution(cmd.Context(), &resources.TaskExecutionRequest{
					Id:    "1",
					Input: "hello",
				})
				if err != nil {
					return errors.Wrap(err, "could not request task execution")
				}

				log.Println("got resp", resp)

				time.Sleep(time.Second)
			}
		},
	}
}
