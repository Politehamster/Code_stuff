package main

import (
	"fmt"
	"_mytemp"
)

func init() {
	rando.Seed()
}

func main() {
	x := make([]int, 10)
	rando.ArrayI(x, 500)
	fmt.Println(x)

}
