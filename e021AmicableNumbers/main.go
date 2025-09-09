/*

Amicable Numbers
Problem 21

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called
amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22,	44, 55 and 110; therefore d(220) = 284. #
The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

Answer:  31626
*/

package main

import "fmt"

func getProperDivisors(n int) []int {
	var divisors []int
	for i := 1; i < n/2+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func getDivisorSum(divisors []int) int {
	var sum int
	for _, divisor := range divisors {
		sum += divisor
	}
	return sum
}

func getAmicableNumbers(n int) []int {
	var amicableNumbers []int
	for i := 1; i < n; i++ {
		divisors := getProperDivisors(i)
		divisorSum := getDivisorSum(divisors)
		if divisorSum == i {
			continue
		}
		divisors = getProperDivisors(divisorSum)
		divisorSum = getDivisorSum(divisors)
		if divisorSum == i {
			amicableNumbers = append(amicableNumbers, i)
		}
	}
	return amicableNumbers
}

func main() {
	amicableNumbers := getAmicableNumbers(10000)
	var sum int
	for _, number := range amicableNumbers {
		sum += number
	}
	fmt.Println(sum)
}
