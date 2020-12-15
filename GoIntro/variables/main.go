package main

import "fmt"

const (
	message = "The answer to life is %d %d\n"
	answer1 = iota
	answer2
)

func main() {
	fmt.Printf(message, answer1, answer2)
}