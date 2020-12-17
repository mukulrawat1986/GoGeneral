package main

import "fmt"

func main() {

	greeting := "Hello, 世界"

	// iterating over the utf-8 encoded string, rune by rune
	for i, r := range greeting {
		fmt.Printf("%d %c\n", i, r)
	}

	fmt.Printf("%s\n", greeting)
}