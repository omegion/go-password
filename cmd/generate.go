package cmd

import (
	"fmt"

	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

// Generate generates password with given params.
func Generate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generates new random password",
		RunE: func(cmd *cobra.Command, args []string) error {
			length, _ := cmd.Flags().GetInt("length")
			numbers, _ := cmd.Flags().GetInt("numbers")
			symbols, _ := cmd.Flags().GetInt("symbols")
			allowUppercase, _ := cmd.Flags().GetBool("allow-uppercase")
			allowRepeat, _ := cmd.Flags().GetBool("allow-repeat")
			custom, _ := cmd.Flags().GetString("custom")

			input := password.GeneratorInput{
				Symbols: custom,
			}

			gen, err := password.NewGenerator(&input)
			if err != nil {
				return err
			}

			res, err := gen.Generate(length, numbers, symbols, !allowUppercase, allowRepeat)
			if err != nil {
				return err
			}

			//nolint:forbidigo // allow to print.
			fmt.Println(res)

			return nil
		},
	}

	setupGetCommand(cmd)

	return cmd
}

// setupAddCommand sets default flags.
func setupGetCommand(cmd *cobra.Command) {
	cmd.Flags().Int("length", 16, "Length of the generated password")
	cmd.Flags().Int("numbers", 4, "Count of numbers")
	cmd.Flags().Int("symbols", 4, "Count of symbols")
	cmd.Flags().Bool("allow-uppercase", true, "Allow uppercase characters")
	cmd.Flags().Bool("allow-repeat", false, "Allow to repeat characters")
	cmd.Flags().String("custom", "", "Custom symbols e.g. \".+=\"")
}
