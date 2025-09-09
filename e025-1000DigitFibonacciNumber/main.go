/*
1000-digit Fibonacci Number
Problem 25

The Fibonacci sequence is defined by the recurrence relation:

	Fn = F{n - 1} + F{n - 2}, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

	F1 = 1
	F2 = 1
	F3 = 2
	F4 = 3
	F5 = 5
	F6 = 8
	F7 = 13
	F8 = 21
	F9 = 34
	F10 = 55
	F11 = 89
	F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

Answer:  4782
*/

package main

import (
	"fmt"
	"math/big"
)

func NextFibunacci(f1, f2 *big.Int) *big.Int {
	return f1.Add(f1, f2)
}

func main() {
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	for i := 3; ; i++ {
		f3 := NextFibunacci(f1, f2)
		f1 = f2
		f2 = f3
		if len(f3.String()) >= 1000 {
			fmt.Println(i)
			break
		}
	}

}
