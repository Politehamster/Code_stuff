package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	for i := range output {
		output[i] = 1
		for k := range output {
			if i != k {
				output[i] *= nums[k]
			}
		}

	}

	return output
}

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(x))
}

/*
func productExceptSelf(nums []int) (output []int) {
    m := make(map[int]int)
    for i, v := range nums {
        if mValue, ok := m[v]; ok {
            output = append(output, mValue)
            continue
        }
        
        m[v] = sum(nums[:i]) * sum(nums[i+1:])
        output = append(output, m[v])
    }
    return output
}

func sum(input []int) int {
    sum := 1
    for _, v := range input {
        if v == 0 {
            return 0
        }
        sum *= v
    }
    return sum
}
*/
/*
Given an array nums of n integers where n > 1,
return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)
*/
