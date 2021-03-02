package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

// Hello function returns a string
func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
