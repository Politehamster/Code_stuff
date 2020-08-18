package main

import "fmt"

func substrBrute(str1, substr1 string) int {
	sublen := len(substr1)

	cnt := 0
	for i := range str1 {
		if str1[i] == substr1[0] {

			if (i + sublen) <= len(str1) {
				for k := i; k < (i + sublen); k++ {
					
					if substr1[cnt] != str1[k] {
						cnt = 0
						break
					}

					cnt++
					if cnt == sublen {
						return i + 1
					}
				}
			}
		}
	}
	return cnt

}

func dfa(substr1 string) {
	
}

func substr_KMP(str1, substr1 string) int{

}



func main() {
	fmt.Println(substrBrute("NaniNOOOO", "NOOO"))

}
