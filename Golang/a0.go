package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 11)
	for i, _ := range a {
		a[i] = rand.Intn(1030)
	}

	fmt.Println(a)
	for i := 0; i < len(a)/2; i++ {
		a[i], a[(len(a)-1)-i] = a[(len(a)-1)-i], a[i]
	}
	fmt.Println(a)
}

/*
package main

import (
	"fmt"
	"math/rand"
)

func sliceRand(a []int) []int {
	for i, _ := range a {
		a[i] = rand.Intn(50)
	}

	fmt.Println(a)
	return a
}

func sliceRev(a []int) []int {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[(len(a)-1)-i] = a[(len(a)-1)-i], a[i]
	}
	fmt.Println(a)
	return a
}

func main() {
	a := make([]int, 11)
	sliceRand(a)
	sliceRev(a)

}
*/

/*
package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]

	}

}

func main() {
	x := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}

	fmt.Println(string(x))
	reverseString(x)
	fmt.Println(string(x))
}
*/

/*
package main

import (
	"fmt"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]

	}

}

func main() {
	k := "MasterClass, © N0 yeah"
	x:=[]byte(k)
	fmt.Println(string(x))
	reverseString(x)
	fmt.Println(string(x))
}

*/

/*???
package main

import (
	"fmt"
)

func reverseString(s string) string {
	x:= make([]byte,len(s))
	k:=0
	for i := len(s) - 1; i >= 0; i-- {
		x[k] += s[i]
		k++
	}
	return string(x)

}

func main() {
	k := "MasterClass, © N0 yeah"

	fmt.Println(k)

	fmt.Println(reverseString(k))
}
*/
