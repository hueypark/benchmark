package slice

import (
	"testing"
)

var (
	singleDimensionSlice []float64
	doubleDimensionSlice [][]float64
)

func init() {
	slice1Len := 1000
	sliec2Len := 2

	singleDimensionSlice = make([]float64, slice1Len*sliec2Len)
	doubleDimensionSlice = make([][]float64, slice1Len)
	for i := 0; i < slice1Len; i++ {
		doubleDimensionSlice[i] = make([]float64, sliec2Len)

	}
}

func BenchmarkClearSlngleDimensionSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clearSlngleDimensionSlice(singleDimensionSlice)
	}
}

func BenchmarkClearDoubleDimensionSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clearDoubleDimensionSlice(doubleDimensionSlice)
	}
}

func clearSlngleDimensionSlice(s []float64) {
	for i := range s {
		s[i] = 0
	}
}

func clearDoubleDimensionSlice(s [][]float64) {
	for i := range s {
		for j := range s[i] {
			s[i][j] = 0
		}
	}
}
