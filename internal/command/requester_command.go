package command

import (
	"github.com/sorenvonsarvort/velas-sphere/internal/service"
	"github.com/spf13/cobra"
)

func NewRequesterCommand() *cobra.Command {
	return &cobra.Command{
		Use: "requester",
		RunE: func(cmd *cobra.Command, args []string) error {
			requester := service.NewRequesterService(
				service.RequesterServiceConfig{
					Target: "127.0.0.1:8081",
				},
			)

			return requester(cmd.Context())
		},
	}
}
