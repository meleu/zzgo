package random

import (
	"math/rand"
	"time"
)

var Seed = time.Now().UnixNano()

func RandomInt(min, max int) int {
	r := rand.New(rand.NewSource(Seed))
	return r.Intn(max-min+1) + min
}
