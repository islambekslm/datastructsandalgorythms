package binarySearch

import "fmt"

func Search(data []int, target int) (int, error) {
	l := 0
	h := len(data) - 1

	for l <= h {
		m := l + (h-l)/2

		if data[m] == target {
			return m, nil
		}

		if data[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return 0, fmt.Errorf("not found")
}

func SearchRecursive(data []int, target int) (int, error) {
	l := 0
	h := len(data) - 1
	m := l + (h-l)/2

	if data[m] == target {
		return m, nil
	}

	if data[m] < target {
		return SearchRecursive(data[m+1:], target)
	}

	if data[m] > target {
		Search(data[:m], target)
	}

	return 0, fmt.Errorf("not found")
}
