package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	// T(v) converts the value v to the type T.
	fmt.Println(x, y, z)

	fmt.Printf("Type %T of %v\n", x, x)
	fmt.Printf("Type %T of %v\n", y, y)
	fmt.Printf("Type %T of %v\n", f, f)
	fmt.Printf("Type %T of %v\n", z, z)
}
