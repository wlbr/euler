/*
Largest prime factor
Problem 3
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Answer: 6857

*/

package main

import (
	"fmt"
	"math"
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
	target := 600851475143
	t := int(math.Trunc(math.Sqrt(float64(target))))

	maxprim := 0
	prims := eras(t)
	for i := 1; i < len(prims); i++ {
		if prims[i] == false && target%i == 0 {
			maxprim = i
		}
	}
	fmt.Println(maxprim)
}
