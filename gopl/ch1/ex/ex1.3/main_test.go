package main

import (
	"bytes"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		echo1(&buffer, []string{"Hello", "World"})
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		echo2(&buffer, []string{"Hello", "World"})
	}
}