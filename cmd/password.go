package cmd

import (
	"fmt"
	"strconv"

	"github.com/meleu/zzgo/pkg/random"
	"github.com/spf13/cobra"
)

const defaultLength = 8

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password [LENGTH]",
	Short: "generates a random password",
	Long:  "generates a random password, with LENGTH characters",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runPassword,
}

func init() {
	rootCmd.AddCommand(passwordCmd)
}

func runPassword(cmd *cobra.Command, args []string) error {
	passwordLength, err := parseLengthArg(args)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), random.New().Password(passwordLength))
	return nil
}

func parseLengthArg(args []string) (int, error) {
	if len(args) == 0 {
		return defaultLength, nil
	}
	return strconv.Atoi(args[0])
}
