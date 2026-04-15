package random

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func SetSeed(seed int64) {
	rng = rand.New(rand.NewSource(seed))
}

func RandomInt(number1, number2 int) int {
	minVal := number1
	maxVal := number2

	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}

	return rng.Intn(maxVal-minVal+1) + minVal
}
