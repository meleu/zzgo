package random

import (
	"math/rand"
	"time"
)

var Seed = time.Now().UnixNano()

func RandomInt(number1, number2 int) int {
	minVal := number1
	maxVal := number2

	if minVal > maxVal {
		minVal, maxVal = maxVal, minVal
	}

	r := rand.New(rand.NewSource(Seed))
	return r.Intn(maxVal-minVal+1) + minVal
}
