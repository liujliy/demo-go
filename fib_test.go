package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < 1; n++ {
		Fib(n)
	}
}
