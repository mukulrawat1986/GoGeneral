package main

import "fmt"

func printer(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	if err := printer("Hello, World!"); err != nil {

	}
}