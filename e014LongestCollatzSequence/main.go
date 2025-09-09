/*

Longest Collatz sequence
Problem 14
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.


Answer: 837799

*/

package main

import (
	"fmt"
)

func even(number int) bool {
	return number%2 == 0
}

func odd(number int) bool {
	return !even(number)
}

func collatz(n int) int {
	if even(n) {
		return n / 2
	}
	return 3*n + 1
}

func collatzsequence(start int) (result []int) {
	c := collatz(start)
	if c == 1 {
		return []int{1, start}
	}
	return append(collatzsequence(c), start)
}

func main() {
	maxl, maxn := 0, 0
	for i := 1; i < 1000000; i++ {
		c := collatzsequence(i)
		if len(c) > maxl {
			maxl = len(c)
			maxn = i
			//fmt.Printf("New Max: %d - %d\n", maxn, maxl)
		}
	}
	fmt.Println(maxn)
}
