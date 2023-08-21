package searchalgorithms

import "fmt"

func BinarySearch(array [10]int, key int) (*int, error) {
	var low, high, mid int
	high = len(array) - 1

	for low <= high {
		mid = (low + high) / 2

		if key == array[mid] {
			return &mid, nil
		} else if key < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}

	err := fmt.Errorf("key not found: %d", key)
	return nil, err
}