package server

import (
	"context"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type PluginServer struct{}

func (p PluginServer) RequestJobExecution(ctx context.Context, req *resources.JobExecutionRequest) (*resources.JobExecutionResponse, error) {
	return &resources.JobExecutionResponse{
		Id:     req.GetId(),
		Output: "world",
	}, nil
}
