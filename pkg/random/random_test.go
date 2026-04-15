package random_test

import (
	"testing"

	"github.com/meleu/zzgo/pkg/random"
)

func TestRandomInt(t *testing.T) {
	t.Parallel()
	random.SetSeed(1)
	got := random.RandomInt(0, 32767)

	// I know it's 545 because I ran it with Seed = 1, and checked the result.
	want := 545
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRandomIntRespectsInterval(t *testing.T) {
	t.Parallel()
	minVal := 5
	maxVal := 10

	for range 50 {
		got := random.RandomInt(minVal, maxVal)
		if got < minVal || got > maxVal {
			t.Errorf("got %d, want a value %d <= n <= %d", got, minVal, maxVal)
		}
	}
}

func TestRandomIntWithEqualBoundsReturnsTheGivenValue(t *testing.T) {
	t.Parallel()
	value := 50
	got := random.RandomInt(value, value)

	if got != value {
		t.Errorf("got %d, want %d", got, value)
	}
}

func TestRandomIntSwapsValuesWhenFirstArgIsGreater(t *testing.T) {
	t.Parallel()
	minVal := 5
	maxVal := 10

	for range 50 {
		got := random.RandomInt(maxVal, minVal)
		if got < minVal || got > maxVal {
			t.Errorf("got %d, expected %d <= n <= %d", got, minVal, maxVal)
		}
	}
}

func TestRandomIntGeneratesDifferentValuesAcrossCalls(t *testing.T) {
	t.Parallel()
	random.SetSeed(1)
	first := random.RandomInt(0, 32767)
	second := random.RandomInt(0, 32767)

	if first == second {
		t.Errorf("consecutive calls returned the same value %d", first)
	}
}
