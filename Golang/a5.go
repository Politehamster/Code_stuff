package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	x := make([]int, 10)
	for c, _ := range x {

		x[c] = rand.Intn(1500)
	}
	
	sort.Ints(x)
	switch {
	case sort.IntsAreSorted(x) == true:
		fmt.Println("Slice sorted, mhm. Look: ", x)
	default:
		fmt.Println("Nani")
	}

	
}
