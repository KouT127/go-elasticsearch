package commands

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	commands := []*cobra.Command{
		newReindexCmd(),
	}

	rootCmd := &cobra.Command{
		Use: "app",
	}
	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}
