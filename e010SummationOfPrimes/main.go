/*

Summation of primes
Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.


Answer: 142913828922
*/

package main

import (
	"fmt"
)

func eras(max int) (nums []bool) {

	nums = make([]bool, max)
	nums[1] = true

	for i := 2; i < max; i++ {
		if nums[i] == false {
			for j := i * i; j < max; j += i {
				nums[j] = true
			}
		}
	}
	return nums
}

func main() {
	result := 0
	prims := eras(2000000)

	for p, isP := range prims {
		if !isP && p != 0 {
			result += p
		}
	}
	fmt.Println(result)
}
