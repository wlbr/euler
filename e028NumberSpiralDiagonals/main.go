/*
Number Spiral Diagonals                                                                                                                                          │
Problem 28                                                                                                                                                       │
Starting with the number 1 and moving to the right in a clockwise direction a                                                                                    │
5 by 5 spiral is formed as follows:                                                                                                                              │

_21_ 22  23  24  _25_                                                                                                                                            │
 20  _7_  8  _9_  10                                                                                                                                             │
 19   6  _1_  2   11                                                                                                                                             │
 18  _5_  4  _3_  12                                                                                                                                             │
_17_ 16  15  14  _13_                                                                                                                                            │
                                                                                                                                                                 │
It can be verified that the sum of the numbers on the diagonals is 101.                                                                                          │
                                                                                                                                                                 │
What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?                                                                 │

Answer: 669171001
*/

package main

import "fmt"

func main() {
	limit := 1001
	sum := 1
	currentNum := 1
	for sideLen := 3; sideLen <= limit; sideLen += 2 {
		step := sideLen - 1
		for i := 0; i < 4; i++ {
			currentNum += step
			sum += currentNum
		}
	}
	fmt.Println(sum)
}
