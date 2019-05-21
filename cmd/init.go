package cmd

import (
	"github.com/spf13/cobra"
)

func initCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize your migraines",
		Run: func(cmd *cobra.Command, args []string) {
			createRunner()
		},
	}
}
