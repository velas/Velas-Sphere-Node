package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"google.golang.org/grpc"
)

type PluginServiceConfig struct {
	ListenPort string `json:"listen_port"`
}

func NewPluginService(config PluginServiceConfig) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		log.Println("plugin service started")

		lis, err := net.Listen("tcp", config.ListenPort)
		if err != nil {
			return fmt.Errorf("failed to listen: %w", err)
		}

		s := grpc.NewServer()
		resources.RegisterPluginServer(s, server.PluginServer{})

		err = s.Serve(lis)
		if err != nil {
			return fmt.Errorf("failed to serve: %w", err)
		}

		return nil
	}
}
