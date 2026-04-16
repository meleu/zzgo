package cmd

import (
	"fmt"
	"strconv"

	"github.com/meleu/zzgo/pkg/random"
	"github.com/spf13/cobra"
)

// randomCmd represents the random subcommand
var randomCmd = &cobra.Command{
	Use:     "random",
	Aliases: []string{"rand"},
	Short:   "zz random - generates a random integer",
	Long: `zz random - generates a random integer

With no arguments, a number between 0 and 32767 (similar to Bash's $RANDOM).
With one argument, a number between 0 and given number.
With two arguments, a number between the given numbers.`,
	Args:         cobra.MaximumNArgs(2),
	SilenceUsage: true,
	RunE:         run,
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func run(cmd *cobra.Command, args []string) error {
	n1 := 0
	n2 := 32767
	var err error

	if len(args) == 1 {
		n2, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid arg %q: must be an integer", args[0])
		}
	} else if len(args) == 2 {
		n1, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid arg %q: must be an integer", args[0])
		}
		n2, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid arg %q: must be an integer", args[1])
		}
	}

	randomNumber := random.New().Int(n1, n2)
	fmt.Fprintln(cmd.OutOrStdout(), randomNumber)

	return nil
}
