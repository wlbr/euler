/*

Number letter counts
Problem 17
If the numbers 1 to 5 are written out in words:
one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words,
how many letters would be used?


NOTE: Do not count spaces or hyphens. For example,
342 (three hundred and forty-two)
contains 23 letters and
115 (one hundred and fifteen)
contains 20 letters.

The use of "and" when writing out numbers is in compliance with British usage.


Answer: 21124

*/

package main

import (
	"fmt"
	"math"
)

func trunc(n int) int {
	return int(math.Trunc(float64(n)))

}

func thousands(n int) (result string, rest int) {
	appendix := " thousand"
	if n%1000 != 0 {
		appendix += " and "
	}
	h := n / 1000
	switch h {
	case 1:
		result = "one" + appendix
	case 2:
		result = "two" + appendix
	case 3:
		result = "three" + appendix
	case 4:
		result = "four" + appendix
	case 5:
		result = "five" + appendix
	case 6:
		result = "six" + appendix
	case 7:
		result = "seven" + appendix
	case 8:
		result = "eight" + appendix
	case 9:
		result = "nine" + appendix
	}
	return result, n % 100
}

func hundreds(n int) (result string, rest int) {
	appendix := " hundred"
	if n%100 != 0 {
		appendix += " and "
	}
	h := n / 100
	switch h {
	case 1:
		result = "one" + appendix
	case 2:
		result = "two" + appendix
	case 3:
		result = "three" + appendix
	case 4:
		result = "four" + appendix
	case 5:
		result = "five" + appendix
	case 6:
		result = "six" + appendix
	case 7:
		result = "seven" + appendix
	case 8:
		result = "eight" + appendix
	case 9:
		result = "nine" + appendix
	}
	return result, n % 100
}

func tens(n int) (result string, rest int) {
	appendix := ""
	if n%10 != 0 {
		appendix = "-"
	}
	t := n / 10
	switch t {
	case 1:
		ti, _ := tensirregular(n % 10)
		return ti, 0
	case 2:
		result = "twenty" + appendix
	case 3:
		result = "thirty" + appendix
	case 4:
		result = "forty" + appendix
	case 5:
		result = "fifty" + appendix
	case 6:
		result = "sixty" + appendix
	case 7:
		result = "seventy" + appendix
	case 8:
		result = "eighty" + appendix
	case 9:
		result = "ninety" + appendix
	}
	return result, n % 10
}

func tensirregular(n int) (result string, rest int) {
	appendix := ""
	switch n {
	case 0:
		result = "ten" + appendix
	case 1:
		result = "eleven" + appendix
	case 2:
		result = "twelve" + appendix
	case 3:
		result = "thirteen" + appendix
	case 4:
		result = "fourteen" + appendix
	case 5:
		result = "fifteen" + appendix
	case 6:
		result = "sixteen" + appendix
	case 7:
		result = "seventeen" + appendix
	case 8:
		result = "eighteen" + appendix
	case 9:
		result = "nineteen" + appendix
	}
	return result, 0
}

func digits(n int) (result string) {
	appendix := ""
	switch n {
	case 1:
		result = "one" + appendix
	case 2:
		result = "two" + appendix
	case 3:
		result = "three" + appendix
	case 4:
		result = "four" + appendix
	case 5:
		result = "five" + appendix
	case 6:
		result = "six" + appendix
	case 7:
		result = "seven" + appendix
	case 8:
		result = "eight" + appendix
	case 9:
		result = "nine" + appendix
		//case 0:
		//	result = "zero" + appendix
	}

	return result
}

func numtostring(n int) (result string) {
	th, r := thousands(n)
	h, r := hundreds(n)
	t, r := tens(r)
	s := digits(r)

	return fmt.Sprintf("%s%s%s%s", th, h, t, s)

}

func speciallen(s string) (result int, ress string) {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '-' {
			result++
			ress += string(s[i])
		}
	}
	return result, ress
}

func main() {
	//t := 115
	//s := numtostring(t)

	//fmt.Println(speciallen(s))
	var res string
	for i := 1; i <= 1000; i++ {
		ns := numtostring(i)
		res += ns
	}
	l, _ := speciallen(res)
	fmt.Println(l)

}
