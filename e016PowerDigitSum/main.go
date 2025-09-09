/*

Power digit sum
Problem 16
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?


Answer: 1366

*/

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func power(n, p int) *big.Int {
	i := new(big.Int)
	prod := new(big.Int)
	prod.SetInt64(1)
	i.SetInt64(int64(n))
	for j := 1; j <= p; j++ {
		prod.Mul(prod, i)
	}
	return prod
}

func addDigits(n *big.Int) (result int) {
	s := n.String()
	for _, d := range s {
		n, _ := strconv.Atoi(string(d))
		result += n
	}
	return result
}

func main() {
	fmt.Println(addDigits(power(2, 1000)))
}
