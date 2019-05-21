package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

func runCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Run your migraines",
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				b    bytes.Buffer
				args = []string{"run"}
			)
			runnerDir := fmt.Sprintf("%s/*.go", migraineDir)
			matches, _ := filepath.Glob(runnerDir)
			for _, match := range matches {
				args = append(args, match)
			}
			c := exec.Command("go", args...)
			c.Stdout = &b
			c.Stderr = &b
			if err := c.Run(); err != nil {
				fmt.Println(fmt.Errorf("migraine run: %v", err))
				fmt.Println(b.String())
				os.Exit(1)
			} else {
				fmt.Println(b.String())
			}
		},
	}
}
