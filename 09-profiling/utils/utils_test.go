package utils

import "testing"

/* func Benchmark_IsPrime_1_97(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_1(97)
	}
}

func Benchmark_IsPrime_2_97(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_2(97)
	}
} */

func Benchmark_GeneratePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes(2, 100)
	}
}
