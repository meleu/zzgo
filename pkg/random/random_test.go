package random_test

import (
	"testing"

	"github.com/meleu/zzgo/pkg/random"
)

func TestGeneratorInt_WithFixedSeed(t *testing.T) {
	t.Parallel()
	rng := random.NewWithCustomSeed(1)
	got := rng.Int(0, 1000)

	// I know it's 881 because I ran it with custom seed = 1
	want := 881
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGeneratorInt_RespectsGivenInterval(t *testing.T) {
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
}

func TestGeneratorInt_WithEqualBoundsReturnsTheGivenValue(t *testing.T) {
	t.Parallel()
	value := 50
	got := random.New().Int(value, value)

	if got != value {
		t.Errorf("got %d, want %d", got, value)
	}
}

func TestGeneratorInt_SwapsValuesWhenFirstArgIsGreater(t *testing.T) {
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
}

func TestGeneratorInt_GeneratesDifferentValuesAcrossCalls(t *testing.T) {
	t.Parallel()
	rng := random.NewWithCustomSeed(1)
	first := rng.Int(0, 32767)
	second := rng.Int(0, 32767)

	if first == second {
		t.Errorf("consecutive calls returned the same value %d", first)
	}
}
