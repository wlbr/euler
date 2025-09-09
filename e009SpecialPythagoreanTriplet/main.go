/*
Special Pythagorean triplet
Problem 9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

Answer: 31875000
*/

package main

import (
	"fmt"
)

func sum(a, b, c int) int {
	return a + b + c
}

func pythtriplettest(a, b, c int) bool {
	return a*a+b*b == c*c
}

func searchTripleProd(target int) (a, b, c int) {
	for i := 1; i < target; i++ {
		for j := 1; j <= target; j++ {
			for k := 1; k <= target; k++ {
				if pythtriplettest(i, j, k) && sum(i, j, k) == target {
					return i, j, k
				}
			}
		}
	}
	return 0, 0, 0
}

func main() {
	t := 1000
	a, b, c := searchTripleProd(t)
	fmt.Println(a * b * c)
}
