package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	k:=len(str)
	for i := range str {
		
		if str[i] >= 'A' && str[i] <= 'Z' {
			str += string(str[i] + 32)

		} else {
			str += string(str[i])
		}
	}
	return str[k:]
}

func main() {
	k := "aaraAAssAW"
	fmt.Println(toLowerCase(k))
}

/*
package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	k:=len(str)
	for i := range str {
		
		if str[i] >= 'A' && str[i] <= 'Z' {
			str += string(str[i] + 32)

		} else {
			str += string(str[i])
		}
	}
	return str[k:]
}

func main() {
	k := "aaraAAssAW"
	k=toLowerCase(k)
	fmt.Println(k)
}

*/