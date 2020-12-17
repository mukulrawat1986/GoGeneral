package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func echo1(w io.Writer, args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintf(w, "%s\n", s)
}

func echo2(w io.Writer, args[]string) {
	s := strings.Join(args, " ")
	fmt.Fprintf(w, "%s\n", s)
}


func main() {
	echo1(os.Stdout, os.Args[1:])
	echo2(os.Stdout, os.Args[1:])
}