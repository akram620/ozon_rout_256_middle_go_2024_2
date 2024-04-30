package main

import "testing"

func Benchmark7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main7()
	}
}
