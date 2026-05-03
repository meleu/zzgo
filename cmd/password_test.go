package cmd

import (
	"bytes"
	"io"
	"strconv"
	"strings"
	"testing"
)

func TestZZPassword(t *testing.T) {
	t.Run("password has 8 characters by default", func(t *testing.T) {
		out := new(bytes.Buffer)
		rootCmd.SetArgs([]string{"password"})
		rootCmd.SetOut(out)
		rootCmd.SetErr(io.Discard)
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Unexpected error! %s", err)
		}

		password := strings.TrimSpace(out.String())
		got := len(password)
		want := 8
		if got != want {
			t.Errorf("expected generated password length to be %d, got %d", want, got)
		}
	})

	t.Run("respect the length given as argument", func(t *testing.T) {
		out := new(bytes.Buffer)
		length := 15
		rootCmd.SetArgs([]string{"password", strconv.Itoa(length)})
		rootCmd.SetOut(out)
		rootCmd.SetErr(io.Discard)

		err := rootCmd.Execute()
		if err != nil {
			t.Errorf("Unexpected error! %s", err)
		}

		password := strings.TrimSpace(out.String())
		got := len(password)
		want := length
		if got != want {
			t.Errorf("expected generated password length to be %d, got %d", want, got)
		}
	})
}
