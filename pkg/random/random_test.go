package random_test

import (
	"regexp"
	"testing"

	"github.com/meleu/zzgo/pkg/random"
)

func TestGeneratorInt(t *testing.T) {
	t.Run("with fixed seed", func(t *testing.T) {
		t.Parallel()
		rng := random.NewWithCustomSeed(1)
		got := rng.Int(0, 1000)

		// I know it's 881 because I ran it with custom seed = 1
		want := 881
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("respects the given interval", func(t *testing.T) {
		t.Parallel()
		minVal := 5
		maxVal := 10
		rng := random.New()

		for range 50 {
			got := rng.Int(minVal, maxVal)
			if got < minVal || got > maxVal {
				t.Errorf("got %d, want a value %d <= n <= %d", got, minVal, maxVal)
			}
		}
	})

	t.Run("with equal bounds returns the given value", func(t *testing.T) {
		t.Parallel()
		value := 50
		got := random.New().Int(value, value)

		if got != value {
			t.Errorf("got %d, want %d", got, value)
		}
	})

	t.Run("swaps values when first arg is greater", func(t *testing.T) {
		t.Parallel()
		minVal := 5
		maxVal := 10
		rng := random.New()

		for range 50 {
			got := rng.Int(maxVal, minVal)
			if got < minVal || got > maxVal {
				t.Errorf("got %d, expected %d <= n <= %d", got, minVal, maxVal)
			}
		}
	})

	t.Run("generates different values accross calls", func(t *testing.T) {
		t.Parallel()
		rng := random.NewWithCustomSeed(1)
		first := rng.Int(0, 32767)
		second := rng.Int(0, 32767)

		if first == second {
			t.Errorf("consecutive calls returned the same value %d", first)
		}
	})
}

func TestGeneratorPassword(t *testing.T) {
	t.Run("generates a password with the specified length", func(t *testing.T) {
		t.Parallel()
		r := random.New()
		want := 8
		password := r.Password(want)
		got := len(password)

		if got != want {
			t.Errorf("expected a password with %d chars, got %d: %s", want, got, password)
		}
	})

	t.Run("generates different values accross calls", func(t *testing.T) {
		t.Parallel()
		g := random.New()
		first := g.Password(10)
		second := g.Password(10)

		if first == second {
			t.Errorf("consecutive calls returned the same password %q", first)
		}
	})
}

func TestGeneratorNumericPassword(t *testing.T) {
	t.Run("generates a password with digits only", func(t *testing.T) {
		t.Parallel()
		g := random.New()
		password := g.NumericPassword(10)
		reDigits := regexp.MustCompile(`^[0-9]+$`)
		if !reDigits.MatchString(password) {
			t.Errorf("expected password to have only digits, got %q", password)
		}
	})
}
