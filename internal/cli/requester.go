package cli

import (
	"github.com/sorenvonsarvort/velas-sphere/internal/requester"
	"github.com/spf13/cobra"
)

func NewRequesterCommand() *cobra.Command {
	return &cobra.Command{
		Use: "requester",
		Run: func(cmd *cobra.Command, args []string) {
			requester.Run(cmd.Context())
		},
	}
}
