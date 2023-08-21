package main

import (
	"fmt"

	datastructures "github.com/eduardor2m/building-foundation/src/data_structures"
	searchalgorithms "github.com/eduardor2m/building-foundation/src/search_algorithms"
)

func main() {
	var opc int

	for opc != 1 {
		fmt.Println("Select an option: ")
		fmt.Println("1. Linear Search")
		fmt.Println("2. Binary Search")
		fmt.Println("3. Linked List")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			array := datastructures.Array
			value, err := searchalgorithms.LinearSearch(array, 5)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Value found at index: ", *value)
			}

		case 2:
			array := datastructures.Array
			value, err := searchalgorithms.BinarySearch(array, 5)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Value found at index: ", *value)
			}

		case 3:
			datastructures.LinkedListExemple()

		default:
			fmt.Println("Invalid option")

		}
	}

}
