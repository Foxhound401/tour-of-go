package main

import "fmt"

// checkout 4.go
// Alternative way to specified same type variable
// x, y share the same int type
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
