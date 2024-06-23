package binarySearch

import "testing"

func BenchmarkBinarySearch_Loop(b *testing.B) {
	data := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		data = append(data, i)
	}
	target := 5

	for i := 0; i < b.N; i++ {
		Search(data, target)
	}
}

func BenchmarkBinarySearch_Recursive(b *testing.B) {
	data := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		data = append(data, i)
	}
	target := 5

	for i := 0; i < b.N; i++ {
		SearchRecursive(data, target)
	}
}
