package main

import "fmt"

func main() {
	// unlike javascript variable declared without explicit initial
	// value are given their zero value.

	// The zero value is:
	// 0 for numeric types.
	// false for boolean.
	// "" (the empty string) for strings
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
