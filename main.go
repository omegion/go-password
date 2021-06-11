package main

import (
	"github.com/omegion/go-cli/cmd"
	"os"

	cmd2 "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:          "go-cli",
		Short:        "Go CLI application template",
		Long:         "Go CLI application template for Go projects.",
		SilenceUsage: true,
	}

	commander := cmd2.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Add(),
			cmd.Subtract(),
		).
		Init()

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}
