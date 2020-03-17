package command

import (
	"github.com/sorenvonsarvort/velas-sphere/internal/service"
	"github.com/spf13/cobra"
)

func NewPluginCommand() *cobra.Command {
	return &cobra.Command{
		Use: "plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			requester := service.NewPluginService(
				service.PluginServiceConfig{
					ListenPort: ":8082",
				},
			)

			return requester(cmd.Context())
		},
	}
}
