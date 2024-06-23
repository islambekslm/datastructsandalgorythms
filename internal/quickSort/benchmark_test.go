package quickSort

import "testing"

var data []int = make([]int, 1000)

func init() {
	for i := 0; i < 1000; i++ {
		data = append(data, i)
	}
}

func BenchmarkQuickSortWithRandomPivot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSortWithRandomPivot(data)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
