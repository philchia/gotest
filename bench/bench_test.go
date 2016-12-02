package bench

import (
	"testing"
)

func TestBenchmark(t *testing.T) {
	type args struct {
		b *testing.B
		f func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{
				&testing.B{N: 100},
				func() {},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Benchmark(tt.args.b, tt.args.f)
		})
	}
}
