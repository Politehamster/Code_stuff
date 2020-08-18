package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	}
func matrixFill(x [][]int) {
	for i := range x {
		for k := range x[i] {
			x[i][k] = rand.Intn(10)
		}
	}
	matrixPrint(x)
}

func matrixPrint(x [][]int) {
	for i := range x {
		fmt.Println(x[i])
	}
	fmt.Println("/////////")
}

func matrixReverse(x [][]int) {
	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-i-1] = x[len(x)-i-1], x[i]
	}
	matrixPrint(x)
}

func matrixRotate(x [][]int) {
	matrixReverse(x)

	for row := range x {
		for col := row; col < len(x[row]); col++ {
			x[row][col], x[col][row] = x[col][row], x[row][col]
		}
	}
}

func matrixEval(x [][]int) [][]int {
	//col
	a := len(x)
	//row
	b := len(x[0])
	//column<row
	switch{
	case a < b: 
		k:=make([]int,b)
		fmt.Println("column < row!")
		for i:=0;i<b-a;i++{
		matrixReverse(x)
		x = append(x,k)
		matrixReverse(x)
		matrixPrint(x)}
		matrixRotate(x)
		
		
	case a>b:
	fmt.Println("LOOL")
	default: matrixRotate(x)	
		
	}
	
	//if finished or quad just go forth yay
	
	return x
}

func main() {
	a := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}

	//matrixFill(a)

	a=matrixEval(a)
	matrixPrint(a)
}

//matrix
