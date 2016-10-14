package generators

import (
	"math/rand"
	"time"
)

type Randomer interface {
	Int31n(int32) int32
	Seed(int64)
}

var globalRandom Randomer

func GetRandx() Randomer {
	if globalRandom == nil {
		seed := time.Now().UTC().UnixNano()
		r := rand.New(rand.NewSource(seed))
		r.Seed(seed)
		SetRandx(r)
	}
	return globalRandom
}

func SetRandx(r Randomer) {
	globalRandom = r
}

//TODO: Create a copy of globalRandom
func NewRandx(seed int64) Randomer {
	newRand := globalRandom
	newRand.Seed(seed)
	return newRand
}
