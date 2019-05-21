package cmd

import (
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	rootCmd := &cobra.Command{
		Use:   "migraine",
		Short: "migraine, a CLI to help running synchronous tasks",
	}

	// Register Command
	rootCmd.AddCommand(initCommand())
	rootCmd.AddCommand(newCommand())
	rootCmd.AddCommand(runCommand())

	// Run Command
	return rootCmd.Execute()
}
