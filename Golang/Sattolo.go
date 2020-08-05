package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
func sattolo(a []int) {

	for k := len(a); k > 1; k-- {
		c := rand.Intn(k - 1)
		a[c], a[k-1] = a[k-1], a[c]
		fmt.Println(c)
		fmt.Println(a)
	}
}

func main() {
	x := make([]int, 10)
	for i := range x {
		x[i] = i + 1
	}
	fmt.Println(x)
	sattolo(x)

}
