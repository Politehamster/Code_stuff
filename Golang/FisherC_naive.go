package main

import (
	"fmt"
	"math/rand"
	"time"
)

var fig = map[int]string{
	2: "Korol'",
	3: "Ladja ",
	4: "Ferz' ",
	5: "Slon' ",
	6: "Kon'  ",
	0: "Empty ",
	1: "Empty ",
}

func initAftereffect(x []int) {

	/*
	   2=king
	   3=ladja
	   4=ferz'
	   5=slon'
	   6=kon'*/
}


func initRandomize(x []int) {
	k := rand.Intn(len(x))
	a := 0
	b,c := false, false
	//King 2
	switch {
	case k > 1:
		x[k] = 2
	case k == 1:
		x[k-1] = 3
		x[k] = 2
		b = true
	case k==len(x):
		x[k-1]=2
		x[k]=3
		c=true
	default:
		x[k+1] = 2
		x[k] = 3
		b = true
	}
	//LeftLadja 3
	if b != true {
		a = rand.Intn(k - 1)
		x[a] = 3
	}
	//rightLadja 3
	a = rand.Intn(len(x))
	if c!=true{
	for a <= k {
		a = rand.Intn(len(x))
	}

	x[a] = 3
}
	//0 Slon' 5
	a = rand.Intn(len(x))
	for x[a] != 0 {
		a = rand.Intn(len(x))
	}
	x[a] = 5
	//kon' 6
	a = rand.Intn(len(x))
	for x[a] > 1 {
		a = rand.Intn(len(x))
	}
	x[a] = 6
	//1 Slon' 5

	a = rand.Intn(len(x))
	for x[a] != 1 {
		a = rand.Intn(len(x))
	}
	x[a] = 5
	//ferz' 4
	a = rand.Intn(len(x))
	for x[a] > 1 {
		a = rand.Intn(len(x))
	}
	x[a] = 4
	
	a = rand.Intn(len(x))
	for x[a] > 1 {
		a = rand.Intn(len(x))
	}
	x[a] = 6
	initAftereffect(x)
}
func initBoard(x []int) {
	for i := range x {
		if i%2 == 0 {
			x[i]++
		}
	}
	initRandomize(x)
}

func main() {
	
	rand.Seed(time.Now().UnixNano())
	x := make([]int, 8)
	for _ = range x {

		fmt.Printf("Peshka ")

	}
	fmt.Println("")
	initBoard(x)
	for _, i := range x {

		fmt.Printf(fig[i])
		fmt.Printf(" ")
	}
}
