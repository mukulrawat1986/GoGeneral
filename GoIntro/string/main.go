package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// take substrings
	fmt.Printf("%s\n", atoz[0:9])

	// go through a string rune by rune
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// only using indices when going through string rune by rune
	for i := range atoz {
		fmt.Printf("%d ", i)
	}

	// get length of string
	fmt.Printf("\nLength of string is: %d\n", len(atoz))
}