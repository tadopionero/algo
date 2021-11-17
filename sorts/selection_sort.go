package main

import "fmt"

func selectionSort(arr []int) {
	size := len(arr)

	for i := 0; i < size-1; i++ {
		min_index := i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}

		temp := arr[min_index]
		arr[min_index] = arr[i]
		arr[i] = temp

	}
	fmt.Println(arr)
}

func test_selection_sort() {
	arr := []int{1, 2, 5, 4, 6, 3}
	selectionSort(arr)
}
