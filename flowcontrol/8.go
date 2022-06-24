// As a simple way to play with functions and loops, implement the quare
// root function using Newton's method.
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < int(x); i++ {
		z -= (math.Pow(z, 2) - x) / (2 * z)
		fmt.Println(i, z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
