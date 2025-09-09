/*

Smallest multiple
Problem 5
2520 is the smallest number that can be divided by each of the
numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible
by all of the numbers from 1 to 20?

Answer 232792560
*/

package main

import (
	"fmt"
)

func divisablityCheck(i, max int) (res bool) {
	res = true

	for j := 2; j < max; j++ {
		if i%j != 0 {
			res = false
		}
	}
	return res
}

func main() {
	max := 20
	for i := 3; true; i++ {
		if divisablityCheck(i, max) {
			fmt.Println(i)
			break
		}
	}
}
