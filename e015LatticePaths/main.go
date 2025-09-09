/*

Lattice paths
Problem 15
Starting in the top left corner of a 2×2 grid, and only being able to move
to the right and down, there are exactly 6 routes to the bottom right corner.

https://projecteuler.net/project/images/p015.gif

How many such routes are there through a 20×20 grid?

Answer: 137846528820

*/

package main

import (
	"fmt"
	"math/big"
)

//The number of lattice paths from  (0,0) to  (n,k) is equal to the binomial coefficient  {n+k} above {n}

func faculty(n int) *big.Int {
	result := new(big.Int)
	result.SetInt64(1)
	for i := 1; i <= n; i++ {
		n := new(big.Int)
		n.SetInt64(int64(i))
		result.Mul(result, n)
	}
	return result
}

func bincoeff(above, under int) *big.Int {
	fa := faculty(above)
	fu := faculty(under)
	fau := faculty(above - under)
	res := fa.Div(fa, fu.Mul(fu, fau))
	return res
}

func main() {
	dimx := 20
	dimy := 20

	fmt.Println(bincoeff(dimx+dimy, dimy))

}
