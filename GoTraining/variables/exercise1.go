package main

import (
	"fmt"
)

func main() {
	// Declare variables that are set to their zero value
	var a int
	var b string
	var c bool

	// Display the value of the variables
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 12
	bb := "hello"
	cc := true

	// Display the value of those variables
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	// Perform a type conversion
	pi := float32(3.14)

	// Display the value of pi
	fmt.Println(pi)
}
