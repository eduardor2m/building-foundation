package searchalgorithms

import (
	"fmt"
)

func LinearSearch(array [10]int, key int) (*int, error) {
	for i := 0; i < len(array); i++ {
		if array[i] == key {
			return &i, nil
		}
	}

	err := fmt.Errorf("key not found: %d", key)
	return nil, err
}
