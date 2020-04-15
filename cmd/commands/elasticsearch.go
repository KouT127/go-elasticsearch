package commands

import "github.com/spf13/cobra"

func newReindexCmd() *cobra.Command {
	rebuildIndexCmd := &cobra.Command{
		Use: "reindex",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return rebuildIndexCmd
}
