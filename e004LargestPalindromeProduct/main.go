/*

Largest palindrome product
Problem 4
A palindromic number reads the same both ways. The largest palindrome
made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609
*/

package main

import (
	"fmt"
	"strconv"
)

func sreverse(s string) string {
	res := make([]rune, len(s))
	for i, c := range s {
		res[len(s)-i-1] = c
	}
	return string(res)
}

func checkPalim(istr string) bool {
	rstr := sreverse(istr)
	return istr == rstr
}

func checkPalimInt(i int) bool {
	istr := strconv.Itoa(i)
	return checkPalim(istr)
}

func findMaxPalindromMadeOfTwoThreeDigitNumbers() (res int) {
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			t := i * j
			if checkPalimInt(t) {
				if t > res {
					res = t
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findMaxPalindromMadeOfTwoThreeDigitNumbers())
}
