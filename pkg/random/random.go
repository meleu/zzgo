// Package random provides a random number generator.
//
// Create a generator (seeded with the current time):
//
//	g := random.New()
//	n := g.Int(1, 10) // random int in [1, 10]
//
// Create a generator with a fixed seed for reproducible output:
//
//	g := random.NewWithCustomSeed(42)
//	n := g.Int(-5, 5) // random int in [-5, 5]
//
// The bounds passed to Int may be given in any order; Int(10, 1) is
// equivalent to Int(1, 10).
//
// Inspired by:
// https://github.com/funcoeszz/funcoeszz/blob/master/zz/zzaleatorio.sh
package random

import (
	"math/rand"
	"time"
)

// Generator is a seedable source of pseudo-random numbers. It is not
// safe for concurrent use.
type Generator struct {
	rng *rand.Rand
}

// New returns a Generator seeded with the current time.
func New() *Generator {
	return &Generator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// NewWithCustomSeed returns a Generator seeded with the given value,
// producing a reproducible sequence of numbers.
func NewWithCustomSeed(seed int64) *Generator {
	return &Generator{
		rng: rand.New(rand.NewSource(seed)),
	}
}

// Int returns a pseudo-random integer in the inclusive range bounded by
// number1 and number2. The bounds may be provided in either order.
func (g *Generator) Int(number1, number2 int) int {
	minVal := number1
	maxVal := number2

	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}

	return g.rng.Intn(maxVal-minVal+1) + minVal
}
