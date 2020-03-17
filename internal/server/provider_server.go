package server

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type ProviderServer struct {
	PluginClient resources.PluginClient
}

func (p ProviderServer) RequestTaskExecution(ctx context.Context, req *resources.TaskExecutionRequest) (*resources.TaskExecutionRequestResponse, error) {
	resp, err := p.PluginClient.RequestJobExecution(ctx, &resources.JobExecutionRequest{
		Id:    req.GetId(),
		Input: req.GetInput(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "could not request job execution from plugin")
	}

	log.Println("got resp", resp)

	return &resources.TaskExecutionRequestResponse{
		Id:     resp.GetId(),
		Output: resp.GetOutput(),
	}, nil
}
