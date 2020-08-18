package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var subStr string
	cnt, max := 0, 0
	for i := 0; i < len(s)-1; i++ {
		cnt++
		subStr+=string(s[i])
		
		for k := 0; k < cnt; k++ {
			if s[k] == s[i+1] {
				if max < cnt {
					max = cnt
				}

				cnt = 0
				fmt.Println(subStr, "\n")
				subStr=""
			}

		}

	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("Helohjhjhihihihihihijgygyppprospplaygroundgyufrtyxsroooezsxrtedplaygroundgyufrtyxsrooo"))
}
