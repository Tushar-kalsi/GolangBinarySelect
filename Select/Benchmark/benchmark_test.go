package Select

import (
	"Select/Select"
	"testing"
)

// BenchmarkSelectU64 benchmarks the select_u64 function.
func BenchmarkSelectU64(b *testing.B) {
	x := uint64(0b0100110101101)
	i := uint64(3)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Select.Selectu64(x, i)
	}
}

// BenchmarkTrailingZeros64 benchmarks the TrailingZeros64 function.
func BenchmarkTrailingZeros64(b *testing.B) {
	x := uint64(0b0000000001000)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		Select.TrailingZeros64(x)
	}
}
