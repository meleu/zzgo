package cmd

import (
	"io"
	"testing"
)

func TestZZRandom_ArgsValidation(t *testing.T) {
	argTests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no args ok",
			args:    []string{"random"},
			wantErr: false,
		},
		{
			name:    "one arg ok",
			args:    []string{"random", "10"},
			wantErr: false,
		},
		{
			name:    "two args ok",
			args:    []string{"random", "10", "1"},
			wantErr: false,
		},
		{
			name:    "three args fails",
			args:    []string{"random", "1", "5", "10"},
			wantErr: true,
		},
		{
			name:    "invalid arg fails",
			args:    []string{"random", "abc"},
			wantErr: true,
		},
		{
			name:    "invalid second arg fails",
			args:    []string{"random", "1", "abc"},
			wantErr: true,
		},
		{
			name:    "float arg fails",
			args:    []string{"random", "5.5"},
			wantErr: true,
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
