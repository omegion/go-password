package cmd

import (
	"fmt"
	"github.com/omegion/go-password/internal/controller"
	"github.com/spf13/cobra"
)

// Generate prints version/build.
func Generate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates new random password",
		RunE: func(cmd *cobra.Command, args []string) error {
			length, _ := cmd.Flags().GetInt("length")
			numbers, _ := cmd.Flags().GetBool("numbers")
			uppercase, _ := cmd.Flags().GetBool("uppercase")
			lowercase, _ := cmd.Flags().GetBool("lowercase")
			symbols, _ := cmd.Flags().GetBool("symbols")
			custom, _ := cmd.Flags().GetString("custom")

			gen := controller.NewGenerator()
			gen.Length = &length
			gen.Numbers = numbers
			gen.Uppercase = uppercase
			gen.Lowercase = lowercase
			gen.Symbols = symbols

			if custom != "" {
				gen.ExtraSymbols = custom
			}

			str, err := gen.Generate()
			if err != nil {
				return err
			}

			fmt.Println(str)
			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().Int("length", 8, "Vault Address")
	cmd.Flags().Bool("numbers", true, "Vault Address")
	cmd.Flags().Bool("uppercase", true, "Vault Address")
	cmd.Flags().Bool("lowercase", true, "Vault Address")
	cmd.Flags().Bool("symbols", true, "Vault Address")
	cmd.Flags().String("custom", "", "Vault Address")
}
