package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)
func init(){
rand.Seed(time.Now().UTC().UnixNano())
}
func BubbleSort(items []int){

for i:=range items{
		for j := 0; j < (len(items) - 1 - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	
}

func main() {
	start := time.Now()
	x := make([]int, 100)
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(200)
		fmt.Println(x[i])
	}
	fmt.Println(x)
	BubbleSort(x)
	fmt.Println(x)
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)

} 