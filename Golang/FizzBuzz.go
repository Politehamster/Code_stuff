/*
Approach 1: Naive Approach

Intuition

The moment you hear of FizzBuzz you think whether the number is divisible by 3, 5 or both.

Algorithm

    Initialize an empty answer list.
    Iterate on the numbers from 1...N1 ... N1...N.
    For every number, if it is divisible by both 3 and 5, add FizzBuzz to the answer list.
    Else, Check if the number is divisible by 3, add Fizz.
    Else, Check if the number is divisible by 5, add Buzz.
    Else, add the number.


Complexity Analysis

    Time Complexity: O(N)O(N)O(N)
    Space Complexity: O(1)O(1)O(1)

Approach 2: String Concatenation

Intuition

This approach won't reduce the asymptotic complexity, but proves to be a neater solution when FizzBuzz comes with a twist. What if FizzBuzz is now FizzBuzzJazz i.e.

3 ---> "Fizz" , 5 ---> "Buzz", 7 ---> "Jazz"

If you try to solve this with the previous approach the program would have too many conditions to check:

    Divisible by 3
    Divisible by 5
    Divisible by 7
    Divisible by 3 and 5
    Divisible by 3 and 7
    Divisible by 7 and 3
    Divisible by 3 and 5 and 7
    Not divisible by 3 or 5 or 7.

This way if the FizzBuzz mappings increase, the conditions would grow exponentially in your program.

Algorithm

Instead of checking for every combination of these conditions, check for divisibility by given numbers i.e. 3, 5 as given in the problem. If the number is divisible, concatenate the corresponding string mapping Fizz or Buzz to the current answer string.

For eg. If we are checking for the number 15, the steps would be:

Condition 1: 15 % 3 == 0 , num_ans_str = "Fizz"
Condition 2: 15 % 5 == 0 , num_ans_str += "Buzz"
=> num_ans_str = "FizzBuzz"

So for FizzBuzz we just check for two conditions instead of three conditions as in the first approach.

Similarly, for FizzBuzzJazz now we would just have three conditions to check for divisibility.

Complexity Analysis

    Time Complexity: O(N)O(N)O(N)
    Space Complexity: O(1)O(1)O(1)

Approach 3: Hash it!

Intuition

This approach is an optimization over approach 2. When the number of mappings are limited, approach 2 looks good. But what if you face a tricky interviewer and he decides to add too many mappings?

Having a condition for every mapping is not feasible or may be we can say the code might get ugly and tough to maintain.

What if tomorrow we have to change a mapping or may be delete a mapping? Are we going to change the code every time we have a modification in the mappings?

We don't have to. We can put all these mappings in a Hash Table.

Algorithm

    Put all the mappings in a hash table. The hash table fizzBuzzHash would look something like { 3: 'Fizz', 5: 'Buzz' }
    Iterate on the numbers from 1...N1 ... N1...N.
    For every number, iterate over the fizzBuzzHash keys and check for divisibility.
    If the number is divisible by the key, concatenate the corresponding hash value to the answer string for current number. We do this for every entry in the hash table.
    Add the answer string to the answer list.

    This way you can add/delete mappings to/from to the hash table and not worry about changing the code.

So, for FizzBuzzJazz the hash table would look something like { 3: 'Fizz', 5: 'Buzz', 7: 'Jazz' }

Complexity Analysis

    Time Complexity : O(N)O(N)O(N)
    Space Complexity : O(1)O(1)O(1)
*/

package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	x := make([]string, n)
	for i := range x {

		switch {
		case i%3 == 2 && i%5 == 4:
			x[i] = "FizzBuzz"
		case i%3 == 2:
			x[i] = "Fizz"
		case i%5 == 4:
			x[i] = "Buzz"
		default:
			x[i] = strconv.Itoa(i + 1)
		}
	}
	return x
}

func main() {
	n := 50
	fmt.Println(fizzBuzz(n))
}

/*Write a program that outputs the string representation of numbers from 1 to n.
But for multiples of three it should output Fizz instead of the number and for the multiples of five output Buzz. For numbers which are multiples of both three and five output FizzBuzz.
Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/
