package cmd

import (
	"bytes"
	"io"
	"slices"
	"strings"
	"testing"
)

func TestZZFlipcoin_OutputIsHeadsOrTails(t *testing.T) {
	validResults := []string{"heads", "tails"}
	seenResults := make(map[string]int, len(validResults))
	iterations := 100

	for range iterations {
		buf := new(bytes.Buffer)
		rootCmd.SetArgs([]string{"flipcoin"})
		rootCmd.SetOut(buf)
		rootCmd.SetErr(io.Discard)
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Unexpected error! %s", err)
		}

		got := strings.TrimSpace(buf.String())
		if !slices.Contains(validResults, got) {
			t.Fatalf("got %q, but expected it to be one of %v", got, validResults)
		}
		seenResults[got]++
	}

	for _, face := range validResults {
		if seenResults[face] == 0 {
			t.Errorf("After %d iterations, %q never appeared!", iterations, face)
		}
	}
}
