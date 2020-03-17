package service

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

type RequesterServiceConfig struct {
	Target string
}

func NewRequesterService(config RequesterServiceConfig) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		log.Println("requester started")

		conn, err := grpc.Dial(config.Target, grpc.WithInsecure())
		if err != nil {
			return errors.Wrap(err, "failed to dial the provider")
		}
		defer conn.Close()

		c := resources.NewProviderClient(conn)

		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			resp, err := c.RequestTaskExecution(ctx, &resources.TaskExecutionRequest{
				Id:    "1",
				Input: "hello",
			})
			if err != nil {
				return errors.Wrap(err, "could not request task execution")
			}

			log.Println("got resp", resp)

			time.Sleep(time.Second)
		}
	}
}
