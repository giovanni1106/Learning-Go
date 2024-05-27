package createSlice

import (
	"testing"
)

func BenchmarkAppendPreallocatedSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Prealocação com capacidade
		sl := make([]int, 10, 20)
		sl = append(sl, i)
	}
}

func BenchmarkAppendExistingSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Sem prealocação explícita
		sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		sl = append(sl, i)
	}
}
