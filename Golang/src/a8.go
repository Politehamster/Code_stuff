/*

    [input] integer nCols

    An integer, the number of theater's columns.

    Constraints: 1 ? nCols ? 1000.

    [input] integer nRows

    An integer, the number of theater's rows.

    Constraints: 1 ? nRows ? 1000.

    [input] integer col

    An integer, the column number of your own seat (with the rightmost column having index 1).

    Constraints: 1 ? col ? nCols.

    [input] integer row

    An integer, the row number of your own seat (with the front row having index 1).

    Constraints: 1 ? row ? nRows.

    [output] an integer

    The number of people who sit strictly behind you and in your column or to the left.
*/



package main

import (
	"fmt"
)

func scanner(c int) int {
	var a float64
	fmt.Print("(1-",c,"): ")
	for i := 0; i != 1; {
		fmt.Scanf("%f\n", &a)
		switch {
		case a >= 1 && a <= float64(c):
			i = 1
		case a > float64(c):
			fmt.Println("Too big. Try again. Max is", c, ".")
		case a == 0:
			fmt.Println("Too small. Try again.")
		default:
			fmt.Println("Numbers only.")	
		
		}
	}
	fmt.Println("")
	return int(a)
}

func main() {
	var (
		nCols = 16
		nRows = 11
		mCol  = 5
		mRow  = 3
		out   = 0
	)
	fmt.Print("Enter number of columns")
	nCols = scanner(1000)

	fmt.Print("Enter number of rows")
	nRows = scanner(1000)
	fmt.Print("Enter My Row")

	mRow = scanner(nRows)
	fmt.Print("Enter My Column")
	mCol = scanner(nCols)
	out = (nRows - mRow) * (nCols - mCol + 1)
	switch {
	case out == 0:
		fmt.Println("Noone! Phew!I'm really lucky! Time to skidaddle.")
	case out < out/5:
		fmt.Println("Only", out, ". I can live with that.")
	default:
		fmt.Println("Oh my.", out, "peeps. It's going be really awkward.")

	}

}

