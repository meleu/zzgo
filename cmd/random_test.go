package cmd

import (
	"io"
	"testing"
)

func TestZZRandom_Arity(t *testing.T) {
	argTests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			"no args ok",
			[]string{"random"},
			false,
		},
		{
			"one arg ok",
			[]string{"random", "10"},
			false,
		},
		{
			"two args ok",
			[]string{"random", "10", "1"},
			false,
		},
		{
			"three args fails",
			[]string{"random", "1", "5", "10"},
			true,
		},
	}

	for _, tt := range argTests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args)
			rootCmd.SetOut(io.Discard)
			rootCmd.SetErr(io.Discard)
			gotErr := rootCmd.Execute()

			if (gotErr != nil) != tt.wantErr {
				t.Errorf(
					"test: %q, args: %v, wantErr: %v, gotErr: %v",
					tt.name,
					tt.args,
					tt.wantErr,
					gotErr,
				)
			}
		})
	}
}
