/*Approach 1: List operation

Algorithm

    Iterate over all the elements in nums\text{nums}nums
    If some number in nums\text{nums}nums is new to array, append it
    If some number is already in the array, remove it

Complexity Analysis

    Time complexity : O(n2)O(n^2)O(n2). We iterate through nums\text{nums}nums, taking O(n)O(n)O(n) time. We search the whole list to find whether there is duplicate number, taking O(n)O(n)O(n) time. Because search is in the for loop, so we have to multiply both time complexities which is O(n2)O(n^2)O(n2).

    Space complexity : O(n)O(n)O(n). We need a list of size nnn to contain elements in nums\text{nums}nums.

Approach 2: Hash Table

Algorithm

We use hash table to avoid the O(n)O(n)O(n) time required for searching the elements.

    Iterate through all elements in nums\text{nums}nums
    Try if hash_tablehash\_tablehash_table has the key for pop
    If not, set up key/value pair
    In the end, there is only one element in hash_tablehash\_tablehash_table, so use popitem to get it

Complexity Analysis

    Time complexity : O(n?1)=O(n)O(n \cdot 1) = O(n)O(n?1)=O(n). Time complexity of for loop is O(n)O(n)O(n). Time complexity of hash table(dictionary in python) operation pop is O(1)O(1)O(1).

    Space complexity : O(n)O(n)O(n). The space required by hash_tablehash\_tablehash_table is equal to the number of elements in nums\text{nums}nums.

Approach 3: Math

Concept

2?(a+b+c)?(a+a+b+b+c)=c2 * (a + b + c) - (a + a + b + b + c) = c2?(a+b+c)?(a+a+b+b+c)=c

Complexity Analysis

    Time complexity : O(n+n)=O(n)O(n + n) = O(n)O(n+n)=O(n). sum will call next to iterate through nums\text{nums}nums. We can see it as sum(list(i, for i in nums)) which means the time complexity is O(n)O(n)O(n) because of the number of elements(nnn) in nums\text{nums}nums.

    Space complexity : O(n+n)=O(n)O(n + n) = O(n)O(n+n)=O(n). set needs space for the elements in nums

Approach 4: Bit Manipulation

Concept

    If we take XOR of zero and some bit, it will return that bit
        a?0=aa \oplus 0 = aa?0=a
    If we take XOR of two same bits, it will return 0
        a?a=0a \oplus a = 0a?a=0
    a?b?a=(a?a)?b=0?b=ba \oplus b \oplus a = (a \oplus a) \oplus b = 0 \oplus b = ba?b?a=(a?a)?b=0?b=b

So we can XOR all bits together to find the unique number.

Complexity Analysis

    Time complexity : O(n)O(n)O(n). We only iterate through nums\text{nums}nums, so the time complexity is the number of elements in nums\text{nums}nums.

    Space complexity : O(1)O(1)O(1).
*/

package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	m := map[int]int{}
	sum := 0

	for i := range nums {
		if m[nums[i]] == 0 {
			m[nums[i]] = nums[i]
			sum += nums[i]
		} else {
			delete(m, nums[i])
			sum -= nums[i]
		}
	}
	return sum

}

func main() {
	x := []int{4,1,2,1,2}
	fmt.Println(singleNumber(x))
}

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