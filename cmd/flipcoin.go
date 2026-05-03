package cmd

import (
	"fmt"

	"github.com/meleu/zzgo/pkg/random"
	"github.com/spf13/cobra"
)

// flipcoinCmd represents the flipcoin command
var flipcoinCmd = &cobra.Command{
	Use:   "flipcoin",
	Short: "shows 'heads' or 'tails' randomly",
	Args:  cobra.NoArgs,
	RunE:  runFlipcoin,
}

func init() {
	rootCmd.AddCommand(flipcoinCmd)
}

func runFlipcoin(cmd *cobra.Command, args []string) error {
	coinFace := [2]string{"heads", "tails"}

	randomNumber := random.New().Int(0, 1)
	fmt.Fprintln(cmd.OutOrStdout(), coinFace[randomNumber])

	return nil
}
