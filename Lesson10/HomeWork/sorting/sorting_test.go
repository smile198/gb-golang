package sorting

import (
	"math/rand"
	"testing"
)

func generateMap(len int) []int {
	res := make([]int, 0, len)

	for i := 0; i < len; i++ {
		res = append(res, rand.Int())
	}

	return res
}

func BenchmarkSort(b *testing.B) {
	b.StopTimer()
	arr := generateMap(10)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		var testArr []int
		copy(testArr, arr)
		b.StartTimer()

		Sort(testArr)
	}
}
