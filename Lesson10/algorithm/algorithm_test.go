package algorithm

import "testing"

func BenchmarkFibb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibb(30)
	}
}
