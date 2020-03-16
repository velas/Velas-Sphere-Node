package cli

import (
	"github.com/sorenvonsarvort/velas-sphere/internal/provider"
	"github.com/spf13/cobra"
)

func NewProviderCommand() *cobra.Command {
	return &cobra.Command{
		Use: "provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			return provider.Run(cmd.Context())
		},
	}
}
