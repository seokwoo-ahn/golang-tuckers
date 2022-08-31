package testcode

import (
	"golang-tuckers/testcode/libs"
	"testing"
)

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libs.Fibonacci1(30)
	}
}
func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		libs.Fibonacci2(30)
	}
}
