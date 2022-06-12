package main

import "fmt"

// constant cannot be declared using := syntax
const Pi = 3.14

func main() {
	const World = "Viet Nam"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
