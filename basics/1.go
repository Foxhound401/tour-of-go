package main

import (
	"fmt"
	"math/rand"
)

func init() {
	rand.Seed(1000)
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
