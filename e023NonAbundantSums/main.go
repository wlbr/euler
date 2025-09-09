/*

Non-Abundant Sums
Problem 23

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of  would be 1 + 2 + 4 + 7 + 14 = 28, which means that
28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called
abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be
written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that
all integers greater than 28123 can be written as the sum of two abundant numbers. However, this
upper limit cannot be reduced any further by analysis even though it is known that the greatest
number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

Answer:  4179871
*/

package main

import "fmt"

const (
	perfect = iota
	abundant
	deficient
)

type Status int

func (s Status) String() string {
	switch s {
	case perfect:
		return "perfect"
	case abundant:
		return "abundant"
	case deficient:
		return "deficient"
	}
	return ""
}

func getProperDivisors(n int) []int {
	var divisors []int
	for i := 1; i < n/2+1; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func getStatusOfInt(n int) Status {
	divisors := getProperDivisors(n)
	sum := 0
	for _, divisor := range divisors {
		sum += divisor
	}
	if sum == n {
		return perfect
	}
	if sum < n {
		return deficient
	}

	return abundant
}

var nums []Status = []Status{}
var abundantnums []int = []int{}

func main() {
	//get status (perfect deficient, abundant) for all numbers
	for i := 0; i <= 28124; i++ {
		nums = append(nums, getStatusOfInt(i))
		if getStatusOfInt(i) == abundant {
			abundantnums = append(abundantnums, i)
		}
	}
	//write all abundants in a map
	allsums := make(map[int]bool)
	for i := 0; i < len(abundantnums); i++ {
		for j := 0; j < len(abundantnums); j++ {
			allsums[abundantnums[i]+abundantnums[j]] = true
		}
	}

	//add to overall sum if number is not in array of sums of abundants.
	sum := 0
	for i := 1; i < 28123; i++ {
		if allsums[i] == false {
			sum += i
		}
	}
	fmt.Println(sum)

}
