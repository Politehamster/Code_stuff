package main

import "fmt"

func facto(x int) int {
	if x == 0 {
		return 1
	}
	return x * facto(x-1)
}

func main() {
var k int
fmt.Print("Enter number:")
fmt.Scanln(&k)
fmt.Print("Yer factorial:")
	fmt.Println(facto(k))
}
