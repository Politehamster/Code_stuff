package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
//rewrite. 1 weighted /=
//Better? Dunno. Need some testing yo.
func rollMaster(d, s int) []int {
	result := make([]int, d)
	for i := 0; i < d; i++ {
		result[i] = rand.Intn(s-1)
		result[i]++
		}


	return result
}
func roll(d, s int) int {
	x := rollMaster(d, s)
	result := 0
	for _, k := range x {
		result += k

	}
	return result
}

func rollBest(d, s, b int) int {
	x := rollMaster(d, s)
	result := 0
	sort.Ints(x)
	/*switch {
	case sort.IntsAreSorted(x) == true:
		fmt.Println("Slice sorted, mhm. Look: ", x)
	default:
		fmt.Println("Nani")
	}*/
	for i := len(x) - b; i < len(x); i++ {
		result += x[i]

	}

	return result
}

//GURPS Ultra Lite not dnd 5e, mhahaha

type baseStats struct {
	st, dx, iq, ht int
}
type baseExtra struct {
	hp, bd_dice, bd_sides int
}
//Bool or not bool?
type baseSkills struct {
	charisma, warrior, wealthy int
}
type creature struct {
	name   string
	stats  baseStats
	extra  baseExtra
	skills baseSkills
}

func (x *creature) talk() {
	fmt.Println("Hewwo. I am ", x.name)
}
func (x *creature) generate() {
	x.stats.st = roll(1, 4)
	x.stats.dx = roll(1, 4)
	x.stats.iq = roll(1, 4)
	x.stats.ht = roll(1, 4)
	switch {
	case x.stats.st == 1:
		x.extra.hp = 8
		x.extra.bd_dice = 1
		x.extra.bd_sides = 3
	case x.stats.st == 2:
		x.extra.hp = 10
		x.extra.bd_dice = 1
		x.extra.bd_sides = 6
	case x.stats.st == 3:
		x.extra.hp = 14
		x.extra.bd_dice = 2
		x.extra.bd_sides = 6
	case x.stats.st == 4:
		x.extra.hp = 18
		x.extra.bd_dice = 3
		x.extra.bd_sides = 6
	}

}
func main(){
x:=new(creature)
x.name="Pikachu"
x.generate()
x.talk()
fmt.Println(x)
}