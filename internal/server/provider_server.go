package server

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

type ProviderServer struct {
	PluginTarget string
}

func (p ProviderServer) RequestTaskExecution(ctx context.Context, req *resources.TaskExecutionRequest) (*resources.TaskExecutionRequestResponse, error) {
	conn, err := grpc.Dial(p.PluginTarget, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial the plugin")
	}
	defer conn.Close()

	c := resources.NewPluginClient(conn)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	resp, err := c.RequestJobExecution(ctx, &resources.JobExecutionRequest{
		Id:    req.GetId(),
		Input: req.GetInput(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not request task execution")
	}

	log.Println("got resp", resp)

	return &resources.TaskExecutionRequestResponse{
		Id:     resp.GetId(),
		Output: resp.GetOutput(),
	}, nil
}
