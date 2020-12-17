package main

import (
	"fmt"
	"os"
)

func main() {
	for i, r := range os.Args {
		fmt.Printf("%d %s\n", i, r)
	}
}