package server

import (
	"context"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type PluginServer struct{}

func (p PluginServer) RequestJobExecution(ctx context.Context, req *resources.JobExecutionRequest) (*resources.JobExecutionRequestResponse, error) {
	return &resources.JobExecutionRequestResponse{
		Id:     req.GetId(),
		Output: "world",
	}, nil
}
