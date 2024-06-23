package quickSort

import "math/rand"

func QuickSortWithRandomPivot(arr []int) []int {
	lenarr := len(arr)
	if lenarr < 2 {
		return arr
	}

	pivotIndex := rand.Intn(lenarr)

	pivot := arr[pivotIndex]
	var lesser []int
	var greater []int
	var equal []int

	for _, v := range arr {
		if v < pivot {
			lesser = append(lesser, v)
		} else if v > pivot {
			greater = append(greater, v)
		} else {
			equal = append(equal, v)
		}
	}

	lesser = QuickSortWithRandomPivot(lesser)
	greater = QuickSortWithRandomPivot(greater)

	return append(append(lesser, equal...), greater...)
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var lesser []int
	var greater []int

	for _, v := range arr[1:] {
		if v < pivot {
			lesser = append(lesser, v)
		} else {
			greater = append(greater, v)
		}
	}

	lesser = QuickSort(lesser)
	greater = QuickSort(greater)

	return append(append(lesser, pivot), greater...)
}
