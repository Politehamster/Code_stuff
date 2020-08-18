package main

import (
	"fmt"
)

const (
	Trillion = 1000000000000
	Billion  = 1000000000
	Million  = 1000000
	Thousand = 1000
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	trillions := num / Trillion
	billions := (num - trillions*Trillion) / Billion
	millions := (num - trillions*Trillion - billions*Billion) / Million
	thousands := (num - trillions*Trillion - billions*Billion - millions*Million) / Thousand
	hundreds := num - trillions*Trillion - billions*Billion - millions*Million - thousands*Thousand

	fmt.Println(trillions, billions, millions, thousands, hundreds)

	result := ""

	if trillions != 0 {
		result += threes(trillions) + " Trillion"
	}
	if billions != 0 {
	
		if result != "" {
			result += " " + threes(billions)
		} else {
			result += threes(billions)
		}
	
		result += " Billion"
	}
	if millions != 0 {
		if result != "" {
			result += " " + threes(millions)
		} else {
			result += threes(millions)
		}
		result += " Million"
	}
	if thousands != 0 {
		if result != "" {
			result += " " + threes(thousands)
		} else {
			result += threes(thousands)
		}
		result += " Thousand"
	}
	if hundreds != 0 {
		if result != "" {
			result += " " + threes(hundreds)
		} else {
			result += threes(hundreds)
		}
	}
	return result
}

func threes(n int) string {
	hundreds := n / 100
	rem := n - hundreds*100

	if hundreds != 0 && rem != 0 {
		return convOnes(hundreds) + " Hundred " + twos(rem)
	}
	if hundreds == 0 && rem != 0 {
		return twos(rem)
	}
	if hundreds != 0 && rem == 0 {
		return convOnes(hundreds) + " Hundred"
	}
	return ""
}

func twos(n int) string {
	if n == 0 {
		return ""
	} else if n < 10 {
		return convOnes(n)
	} else if n < 20 {
		return convTwoSpecial(n)
	} else {
		tenz := n / 10
		rem := n - tenz*10
		res := convTens(tenz)
		if rem != 0 {
			res += " " + convOnes(rem)
		}
		return res
	}
	// unreachable
	return ""
}

var ones = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var twoDigitSpecial = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var tens = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

func convOnes(n int) string {
	if s, ok := ones[n]; ok {
		return s
	}
	return ""
}

func convTwoSpecial(n int) string {
	if s, ok := twoDigitSpecial[n]; ok {
		return s
	}
	return ""
}

func convTens(n int) string {
	if s, ok := tens[n]; ok {
		return s
	}
	return ""
}

func main() {
	fmt.Println(numberToWords(214748364701235))
}
