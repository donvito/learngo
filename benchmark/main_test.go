package main

import (
	"testing"
)

func benchmarkSum(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(i)
	}
}

func BenchmarkSum10(b *testing.B)       { benchmarkSum(10, b) }
func BenchmarkSum100(b *testing.B)      { benchmarkSum(100, b) }
func BenchmarkSum1000(b *testing.B)     { benchmarkSum(1000, b) }
func BenchmarkSum10000(b *testing.B)    { benchmarkSum(10000, b) }
func BenchmarkSum100000(b *testing.B)   { benchmarkSum(100000, b) }
func BenchmarkSum1000000(b *testing.B)  { benchmarkSum(1000000, b) }
func BenchmarkSum10000000(b *testing.B) { benchmarkSum(10000000, b) }
