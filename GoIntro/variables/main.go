package main

import "fmt"

var (
	message = "The answer to life is %d\n"
	answer = 42
)

func main() {
	fmt.Printf(message, answer)
}