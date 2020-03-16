package provider

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

func Run(ctx context.Context) error {
	log.Println("provider started")

	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "failed to dial the injector")
	}
	defer conn.Close()

	c := resources.NewRequesterClient(conn)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resp, err := c.RequestTaskAcceptance(ctx, &resources.TaskAcceptanceRequest{
			Id:     "1",
			Output: "world",
		})
		if err != nil {
			return errors.Wrap(err, "could not add record")
		}

		log.Println("got resp", resp)
	}
}
