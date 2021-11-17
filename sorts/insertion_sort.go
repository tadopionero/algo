package main

import "fmt"

func insertionSort(arr []int) {
	size := len(arr)

	for i := 0; i < size; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
	fmt.Println(arr)
}

func test_insertion_sort() {
	arr := []int{1, 2, 5, 4, 6, 3}
	insertionSort(arr)
}
