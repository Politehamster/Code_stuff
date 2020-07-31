package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func testChew(x int) int {
	defer wg.Done()
	for i := 0; i < x; i++ {
		x += i
	}
	return x
}

func testGo(c chan int, x int) {
	defer wg.Done()
	c <- x + 10
	// fmt.Println(testChew(x))
}

func testWait(c chan){
wg.Wait()
defer close(c)
}
func main() {

	choo := make(chan int, 20)
	for i := 0; i < 10; i++ {
		go testGo(choo, i)
		wg.Add(1)

	}
	
	testWait(c)
	for k := range choo {
		fmt.Println(k)

	}

}

func testNono


type robotrek struct{
name string
var a1 func testNono (s string){
	
}

}

/*
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func mankey(a chan byte, b []byte) {
	defer wg.Done()
	for _, k := range b {
		a <- byte(k)
	}
}

func main() {
	stre := "Hello, playground"
	c := make(chan byte, 201)
	b := make([]byte, len(stre))

	for i, k := range stre {
		b[i] = byte(k)
	}

	for i := 0; i < 5; i++ {

		go mankey(c, b)
		wg.Add(1)
	}

	wg.Wait()
	close(c)

	for i := range c {
		fmt.Println(string(i))

	}
}

*/