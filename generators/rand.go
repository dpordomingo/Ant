package generators

import (
	"math/rand"
	"time"
)

//Randomer is a basic random generator interface (Int31n and Seed)
type Randomer interface {
	Int31n(int32) int32
	Seed(int64)
}

//globalRandom stores the APP Randomer
var globalRandom Randomer

//GetRandx returns the defined APP Randomer
func GetRandx() Randomer {
	if globalRandom == nil {
		seed := time.Now().UTC().UnixNano()
		r := rand.New(rand.NewSource(seed))
		r.Seed(seed)
		SetRandx(r)
	}

	return globalRandom
}

//SetRandx sets the defined APP Randomer
func SetRandx(r Randomer) {
	globalRandom = r
}

//NewRandx returns a new Randomer based on the set APP Randomer
//TODO: Create a copy of globalRandom
func NewRandx(seed int64) Randomer {
	newRand := globalRandom
	newRand.Seed(seed)
	return newRand
}
