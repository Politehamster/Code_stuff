package main

import "fmt"
import "time"

func main() {
	add := func(x int) int {
		for i := 0; i < x; i++ {
			fmt.Println("Let's do something dumb.")
			time.Sleep(550)
			fmt.Println("You may ask what?. ", i, ".")
			
		}
		return x
	}
	for k:=0; k<10; k++{
	go add(5)}
	var input string
	fmt.Scanln(&input)
	fmt.Println(add(1))
}
