/*Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1

Example 2:

Input: [4,1,2,1,2]
Output: 4

*/

package main

import (
	"fmt"
)

/*Approach 1: List operation

Algorithm

    Iterate over all the elements in nums\text{nums}nums
    If some number in nums\text{nums}nums is new to array, append it
    If some number is already in the array, remove it
*/
func singleNumber(x []int) int {
	c := make(map[int]int)
	for i := range x {
		if c[x[i]] == 0 {
			c[x[i]]++
		} else {
			delete(c, x[i])
		}

	}
	for i := range c {
		return i
	}

	return 0
}

func main() {
	a := []int{2, 4, 2, 991, 3, 3, 4}
	fmt.Println("Our array:",a)
	fmt.Println("Look we found our loner!:",singleNumber(a))

}
