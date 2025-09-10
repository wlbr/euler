/*
Quadratic Primes
Problem 27

Euler discovered the remarkable quadratic formula:

	n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive integer
values 0≤n≤39. However, when n=40,402+40+41=40(40+1)+41 is divisible by 41,
and certainly when n=41,41²+41+41 is clearly divisible by 41.

The incredible formula n²−79n+1601 was discovered, which produces 80 primes for the
consecutive values 0≤n≤79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

	n² + an + b, where |a|<1000 and |b|≤1000

where |n| is the modulus/absolute value of n
e.g. |11|=11 and |−4|=4

Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n,
starting with n=0.

Answer:  -59231
*/
package main

import (
	"fmt"
	"math"
)

// isPrime checks if a number is prime.
// A map is used for memoization to speed up repeated checks.
var primeMemo = make(map[int]bool)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if isPrime, ok := primeMemo[n]; ok {
		return isPrime
	}
	if n == 2 {
		primeMemo[n] = true
		return true
	}
	if n%2 == 0 {
		primeMemo[n] = false
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			primeMemo[n] = false
			return false
		}
	}
	primeMemo[n] = true
	return true
}

func main() {
	maxPrimes := 0
	bestA := 0
	bestB := 0

	// b must be prime for n=0, so we can pre-calculate primes up to 1000
	primesB := []int{}
	for i := 2; i <= 1000; i++ {
		if isPrime(i) {
			primesB = append(primesB, i)
		}
	}

	for a := -999; a < 1000; a++ {
		for _, b := range primesB {
			n := 0
			for {
				val := n*n + a*n + b
				if !isPrime(val) {
					break
				}
				n++
			}
			if n > maxPrimes {
				maxPrimes = n
				bestA = a
				bestB = b
			}
		}
	}

	// The problem asks for the product of the coefficients.
	fmt.Println(bestA * bestB)
}
