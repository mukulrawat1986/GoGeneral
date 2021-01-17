package main

import "fmt"

func main() {
	greeting := "こんにちは世界"
	fmt.Println(greeting)
	fmt.Println("Length of greeting = ", len(greeting))
	for i, char := range greeting {
		fmt.Printf("%d %c\n", i, char)
	}
}