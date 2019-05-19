package cmd

import (
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Generate a new migraine",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			createDir()
			createTemplate(args[0])
		},
	}
}
