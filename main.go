package main

import (
	"os"

	cmdr "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/go-password/cmd"
)

func main() {
	root := &cobra.Command{
		Use:          "password",
		Short:        "Password Generator CLI application",
		Long:         "Password Generator CLI application",
		SilenceUsage: true,
	}

	commander := cmdr.NewCommander(root).
		SetCommand(
			cmd.Version(),
			cmd.Generate(),
		).
		Init()

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}
