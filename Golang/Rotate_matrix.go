package main

import (
	"fmt"
)
func rotate(matrix [][]int)  {
    reverse(matrix)
for i:=range matrix{
	fmt.Println(matrix[i])

}
    for row := range matrix {
        for col := row; col<len(matrix[row]);col++ {
            matrix[row][col], matrix[col][row] =  matrix[col][row], matrix[row][col]
        }
    } 
}

func reverse(matrix [][]int) {
    l := len(matrix)
    for i:=0;i<l/2;i++{
        matrix[i], matrix[l-i-1] =  matrix[l-i-1], matrix[i]
    }
}

func main() {
	mat := [][]int{
		{1, 2,   3,  4},
		{5, 6,   7,  8},
		{9, 10, 11, 12},
		{13,14, 15, 16}}
	for i:=range mat{
	fmt.Println(mat[i])

}
fmt.Println("")

	rotate(mat)
	fmt.Println("")
	for i:=range mat{
	fmt.Println(mat[i])

}
}

/*
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
Example 1:

  [1,2,3],||[7,8,9],[7,4,1]		
  [4,5,6],||[4,5,6],[8,5,2]
  [7,8,9],||[1,2,3],[9,6,3]

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
*/
