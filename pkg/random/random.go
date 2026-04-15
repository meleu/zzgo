package random

import (
	"math/rand"
	"time"
)

type Generator struct {
	rng *rand.Rand
}

func New() *Generator {
	return &Generator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func NewWithCustomSeed(seed int64) *Generator {
	return &Generator{
		rng: rand.New(rand.NewSource(seed)),
	}
}

func (g *Generator) Int(number1, number2 int) int {
	minVal := number1
	maxVal := number2

	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}

	return g.rng.Intn(maxVal-minVal+1) + minVal
}
