package main

import "fmt"

func main() {
	nine := uint8(9)
	fmt.Printf("%b\n", nine)

	var isTrue bool
	fmt.Printf("Value: %t\n", isTrue)

	b := byte(87)
	fmt.Printf("Value: %b\n", b)
}