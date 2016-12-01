package bench

import (
	"testing"
)

func Benchmark(b *testing.B, f func()) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f()
	}
}
