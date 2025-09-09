/*

Lexicographic Permutations
Problem 24

A permutation is an ordered arrangement of objects. For example,
3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically,
we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

Answer:  2783915460
*/

package main

import (
	"fmt"
	"sort"
)

func mergeArrayToInt(arr []int) int {
	var result int
	var power int = 1
	for i := len(arr) - 1; i >= 0; i-- {
		result += arr[i] * power
		power *= 10
	}
	return result
}

// createAllPermutations generates all permutations of a slice of integers.
// It uses a recursive approach to generate the permutations.
func createAllPermutations(nums []int) []int {
	var result []int
	var generate func(arr []int, n int)

	generate = func(arr []int, n int) {
		if n == 1 {
			// We have a full permutation, add a copy to the result.
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			result = append(result, mergeArrayToInt(tmp))
			return
		}

		for i := 0; i < n; i++ {
			generate(arr, n-1)
			// Swap elements to generate the next permutation.
			if n%2 == 1 {
				arr[0], arr[n-1] = arr[n-1], arr[0]
			} else {
				arr[i], arr[n-1] = arr[n-1], arr[i]
			}
		}
	}

	generate(nums, len(nums))
	sort.Ints(result)
	return result
}

func main() {
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	permutations := createAllPermutations(digits)

	fmt.Println(permutations[999999])
}
