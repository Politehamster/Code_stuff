package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func QuickSort(items []int) {

	if len(items) > 1 {
		pivot_index := len(items) / 2
		var smaller_items = []int{}
		var larger_items = []int{}

		for i := range items {
			val := items[i]
			if i != pivot_index {
				if val < items[pivot_index] {
					smaller_items = append(smaller_items, val)
				} else {
					larger_items = append(larger_items, val)
				}
			}
		}

		QuickSort(smaller_items)
		QuickSort(larger_items)

		var merged []int = append(append(append([]int{}, smaller_items...), []int{items[pivot_index]}...), larger_items...)
		//merged := MergeLists(smaller_items, items[pivot_index], larger_items)

		for j := 0; j < len(items); j++ {
			items[j] = merged[j]
		}

	}

}

func main() {
	start := time.Now()
	a1 := make([]int, 100)
	for i := range a1{
		a1[i] = rand.Intn(100)
	}
	elapsed := time.Since(start)
	
	start2 := time.Now()
	fmt.Println(a1)
	elapsed2 := time.Since(start2)

	start3 := time.Now()
	QuickSort(a1)
	elapsed3 := time.Since(start3)
	
	start4 := time.Now()	
	fmt.Println(a1)
	elapsed4 := time.Since(start4)

	log.Printf("created %s", elapsed)
	log.Printf("printed %s", elapsed2)

	log.Printf("sorted %s", elapsed3)

	log.Printf("printed sorted %s", elapsed4)
	
}

