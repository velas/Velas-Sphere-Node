package server

import (
	"context"
	"time"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
)

type PluginServer struct{}

func (p PluginServer) RequestJobExecution(ctx context.Context, req *resources.JobExecutionRequest) (*resources.JobExecutionResponse, error) {

	// before := time.Now()
	output := ""
	for _, c := range "world" {
		output = string(append([]rune(output), c))
		time.Sleep(time.Millisecond * 15)
	}
	// diff := time.Now().Sub(before)
	// TODO: measure time
	return &resources.JobExecutionResponse{
		// Took: diff
		Id:     req.GetId(),
		Output: "world",
	}, nil
}
