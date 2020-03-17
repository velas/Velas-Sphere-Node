package command

import (
	"fmt"
	"log"
	"net"

	"github.com/sorenvonsarvort/velas-sphere/internal/resources"
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewPluginCommand() *cobra.Command {
	return &cobra.Command{
		Use: "plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			lis, err := net.Listen("tcp", ":8082")
			if err != nil {
				return fmt.Errorf("failed to listen: %w", err)
			}

			s := grpc.NewServer()
			resources.RegisterPluginServer(s, server.PluginServer{})

			log.Println("plugin service started")

			return s.Serve(lis)
		},
	}
}
