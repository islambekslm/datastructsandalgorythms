package quicksort

import (
	"datastructuresandalgorythms/internal/quickSort"
	"fmt"
)

func Run() {
	data := []int{4, 3, 5, 7, 222, 3, 5, 1, 88, 5, 78, 100000000000000, 2434525, 58463, 99988, 231312, 4353535353535, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1}
	fmt.Println(data)
	data = quickSort.QuickSort(data)
	fmt.Println(data)
}
