package bench

import (
	"testing"
)

// Benchmark will measure f
func Benchmark(b *testing.B, f func()) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}
