/*
n! means n*(n-1)*...*3*2*1

For example, 10!=10*9*8*7*6*5*4*3*2*1=3628800
and the sum of the digits in the number 10! is 3+6+2+8+8+0+0=27

Find the sum of the digits in the number 100!

*/

package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	fact := big.NewInt(1)
	for i := 2; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	return fact
}

func sumOfDigits(n *big.Int) int {
	sum := 0
	s := n.String()
	for _, digit := range s {
		sum += int(digit - '0')
	}
	return sum
}

func main() {
	n := 100
	fact := factorial(n)
	fmt.Println("100! is:", fact)
	sum := sumOfDigits(fact)
	fmt.Println("Sum of the digits of 100! is:", sum)
}
