package command

import (
	"github.com/sorenvonsarvort/velas-sphere/internal/server"
	"github.com/sorenvonsarvort/velas-sphere/internal/service"
	"github.com/spf13/cobra"
)

func NewProviderCommand() *cobra.Command {
	return &cobra.Command{
		Use: "provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := service.NewProviderService(
				service.ProviderServiceConfig{
					ListenPort: ":8081",
					Server: server.ProviderServer{
						PluginTarget: "127.0.0.1:8082",
					},
				},
			)

			return provider(cmd.Context())
		},
	}
}
