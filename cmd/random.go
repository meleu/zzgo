package cmd

import (
	"fmt"
	"strconv"

	"github.com/meleu/zzgo/pkg/random"
	"github.com/spf13/cobra"
)

// randomCmd represents the random subcommand
var randomCmd = &cobra.Command{
	Use:     "random [number1 [number2]]",
	Args:    cobra.MaximumNArgs(2),
	Aliases: []string{"rand"},
	Short:   "zz random - generates a random integer",
	Long: `zz random - generates a random integer

With no arguments, a number between 0 and 32767 (similar to Bash's $RANDOM).
With one argument, a number between 0 and given number.
With two arguments, a number between the given numbers.`,
	SilenceUsage: true,
	RunE:         run,
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func run(cmd *cobra.Command, args []string) error {
	n1, n2, err := boundsFromArgs(args)
	if err != nil {
		return err
	}

	randomNumber := random.New().Int(n1, n2)
	fmt.Fprintln(cmd.OutOrStdout(), randomNumber)

	return nil
}

// boundsFromArgs parses args into the integers.
//
// Returns an error if any arg is not an integer.
//
// With no args it returns 0 and 32767;
// With one arg, it becomes the upper bound (lower one is 0);
// With two args they are the lower and upper bounds.
func boundsFromArgs(args []string) (n1, n2 int, err error) {
	parsedArgs := make([]int, len(args))
	for i, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			return 0, 0, fmt.Errorf("invalid arg %q: must be an integer", arg)
		}
		parsedArgs[i] = value
	}

	n1, n2 = 0, 32767
	switch len(parsedArgs) {
	case 1:
		n2 = parsedArgs[0]
	case 2:
		n1, n2 = parsedArgs[0], parsedArgs[1]
	}
	// no need to worry about cases where len > 2, it's handled by cobra

	return n1, n2, nil
}
