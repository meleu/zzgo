package random_test

import (
	"testing"

	"github.com/meleu/zzgo/pkg/random"
)

func TestRandomInt(t *testing.T) {
	random.Seed = 1
	got := random.RandomInt(0, 32767)

	// I know it's 545 because I ran it with Seed = 1, and checked the result.
	want := 545
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
