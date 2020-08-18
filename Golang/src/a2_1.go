package main

import (
	"fmt"
)

func foF(x int) (int, int) {
	x += 10
	return x, x / 2
}

func main() {
	k := 0
	_, k = foF(k)
	fmt.Println(k)
}
/*
package main

import (
	"fmt"
)

func foF(x int) (ax int, ay int) {
	ax = x + 10
	ay = ax / 2
	return
}

func main() {
	k := 0
	_, k = foF(k)
	fmt.Println(k)
}
*/

/*
// n! = n?(n-1)! where n >0
func getFactorial(num int) int {
	if num > 1 {
		return num * getFactorial(num-1)
	} else {
		return 1 // 1! == 1
	}
}

Above getFactorial function is recursive, as we are calling getFactorial from inside getFactorial function. The steps to understand are very simple.

When getFactorial get called with a int parameter num, if num is equal to 1, function returns 1, else it goes inside if block and executes num * getFactorial(num-1). Since, there is a function call, function getFactorial called again and return value will be kept on hold until getFactorial returns something. This stack will be kept on building until getFactorial returns something, which is 1eventually. As soon as that happens, all call stack will be resolved one by one eventually resolving first getFactorial call.*/


/*
package main

import "fmt"

var add = func(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("5+3 =", add(5, 3))
}
*/

/*
package main

import "fmt"

func main() {
	sum := func(a int, b int) int {
		return a + b
	}(3, 5)
	
	fmt.Println("5+3 =", sum)
}
*/