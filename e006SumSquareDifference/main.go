/*

Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

Answer: 25164150
*/

package main

import (
	"fmt"
)

func sumOfSquares(max int) (result int) {
	for i := 1; i <= max; i++ {
		result += i * i
	}
	return result
}

func squareOfSum(max int) (result int) {
	for i := 1; i <= max; i++ {
		result += i
	}
	return result * result
}

func main() {
	max := 100
	fmt.Println(squareOfSum(max) - sumOfSquares(max))
}
