/*

10001st prime
Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

Answer:104743
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
	count, result := 0, 0
	max := 10001
	prims := eras(500000)

	for p, isP := range prims {
		if !isP && p != 0 {
			count++
			if count >= max {
				result = p
				break
			}
		}
	}
	fmt.Println(result)
}
