package cmd

import (
	"fmt"
	"strconv"

	"github.com/meleu/zzgo/pkg/random"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:     "random",
	Aliases: []string{"rand"},
	Short:   "zz random - generates a random integer",
	Long: `
With no arguments, generates a number between 0 and 32767 (similar to Bash's $RANDOM).
With one argument, a number between 0 and given number.
With two arguments, a number between the given numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		n1 := 0
		n2 := 32767

		if len(args) == 1 {
			n2, _ = strconv.Atoi(args[0])
		} else if len(args) == 2 {
			n1, _ = strconv.Atoi(args[0])
			n2, _ = strconv.Atoi(args[1])
		}

		randomNumber := random.New().Int(n1, n2)
		fmt.Println(randomNumber)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
