package requester

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"google.golang.org/grpc"
)

type server struct {
}

func (s server) RequestTaskAcceptance(ctx context.Context, req *resources.TaskAcceptanceRequest) (*resources.TaskAcceptanceRequestResponse, error) {
	return &resources.TaskAcceptanceRequestResponse{
		Id: req.GetId(),
	}, nil
}

func Run(ctx context.Context) error {
	log.Println("requester started")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	resources.RegisterRequesterServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
