package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

// Hello function returns a string
func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
