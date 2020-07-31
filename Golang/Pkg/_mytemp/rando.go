package rando

import (
	"math/rand"
	"time"
)

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func ArrayI(x []int, c int) {
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(c)
	}

}
